package main

import (
  "fmt"
  "io/ioutil"
  s "strings"
  "os"
)

var p = fmt.Println
var pp = fmt.Print
var file = os.Args[1]
var content = make([]string, 0)
var struct_name string
var use_db bool
var use_fmt bool
var use_json bool

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func add(new_line string)[]string {
  content = append(content, new_line)
  return content
}

func add_once(new_line string)[]string {
  var exists bool
  for _, line := range content {
    if new_line == line { exists = true }
  }
  if !exists { content = append(content, new_line) }
  return content
}

func check_indent(line string)int {
  chars := []byte(line)
  indent := 0
  for _, char := range chars {
    if char == 32 { indent++ }
  }
  return indent
}

func clear(line string)string {
  return s.Replace(line, " ", "", -1)
}

func build(lines []string, line string, index int, z int, f string) {
  current_indent := check_indent(lines[index])
  if current_indent == 0 {
    if len(line)>0 {
      if s.Contains(line, ":") {
        add(`main["`+s.Split(line, ":")[0]+`"] = "`+s.Split(line, ":")[1]+`"`)
      } else {
        add(`main["`+line+`"] = `+line)
      }
    }
    return
  }
  parent_indent := check_indent(lines[index-z])
  parent := lines[index-z]
  parent = clear(parent)
  line = clear(line)

  if current_indent > parent_indent {
    if f == "n" { add(parent+`["`+line+`"] = `+line) }
    if f == "k" {
      add(parent+`["`+s.Split(line, ":")[0]+`"] = "`+s.Split(line, ":")[1]+`"`)
    }
  } else {
    build(lines, line, index, z+1, f)
  }
}

func init_struct(lines []string) {
  add(`main := map[string]interface{}{}`)
  for index, line := range lines {
    if s.Contains(line, "=") { return }
    if s.Contains(line, ":") {
      build(lines, line, index, 1, "k")
    } else {
      if len(line)>0 { add(clear(line)+` := map[string]interface{}{}`) }
      build(lines, line, index, 1, "n")
      if index == 0 {
        struct_name = line
      }
    }
  }
}

func init_db(lines []string) {
  for _, line := range lines {
    if s.Contains(line, "db") {
      line = s.Replace(line, "db:", "", -1)
      params := s.Split(line, "|")
      add(`db, _ := sql.Open("mysql", "`+params[0]+`:`+params[1]+`@/`+params[2]+`")`)
      add(`defer db.Close()`)
    }
  }
}

func before(lines []string) {
  add(`package main`)
  add(`import (`)

  for _, line := range lines {
    if s.Contains(line, "db") { use_db = true }
    if s.Contains(line, "=") && !use_db { use_json = true }
  }

  if use_db || use_json { use_fmt = true }
  if use_db { add(`  "database/sql"`) }
  if use_db { add(`  _ "github.com/go-sql-driver/mysql"`) }
  if use_json { add(`  "encoding/json"`) }
  if use_fmt { add(`  "fmt"`) }

  add(`)`)

  if !use_db {
    add(`func escape_print(j []byte)[]byte {`)
    add(` return j`)
    add(`}`)
    add(`func unquote(line string)string {`)
    add(`return line[1:len(line)-1]`)
    add(`}`)
  }

  add(`func main() {`)
}

func parse_db_query(lines []string) {
  for _, line := range lines {
    if s.Contains(line, "=") {
      sql_print := s.Split(line, "/")
      table := s.Split(sql_print[0], "=")[1]
      id := sql_print[1]
      field := sql_print[2]
      add_once(`var `+field+` string`)
      add(`db.QueryRow("SELECT `+field+` FROM `+table+` WHERE id=?", `+id+`).Scan(&`+field+`)`)
      add(`fmt.Println(`+field+`)`)
    }
  }
}

func parse_json_query(lines []string) {
  for _, line := range lines {
    if s.Contains(line, "=") {
      value := s.Split(line, "= ")[1]
      if s.Contains(line, "/") {
        parts := s.Split(value, "/")
        object := parts[0]+`["`+parts[1]+`"]`
        add(`respond, _ := json.Marshal(`+object+`)`)
        add(`fmt.Println(unquote(string(respond)))`)
      } else {
        add(`respond, _ := json.Marshal(main["`+value+`"])`)
        add(`respond = escape_print(respond)`)
        add(`fmt.Println(respond)`)
      }
    }
  }
}

func after(lines []string) {
  if use_db == true { parse_db_query(lines) } else { parse_json_query(lines) }
  add(`}`)
}

func write_file() {
  f, _ := os.Create("/tmp/"+file+".go")
  for _, line := range content {
    f.WriteString(line+"\n")
  }
}

func main() {
  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  before(lines)
  if use_db { init_db(lines) } else { init_struct(lines) }
  after(lines)
  write_file()
}
