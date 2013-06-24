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
  //db, _ := sql.Open("mysql", "root:@/go_learn")
  //defer db.Close()

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

func before(lines []string) {
  add(`package main`)
  add(`import (`)

  for _, line := range lines {
    if s.Contains(line, "db") { use_db = true }
    //if s.Contains(line, "=") { use_json = true }
    use_fmt = true
  }

  if use_db { add(`  "database/sql"`) }
  if use_db { add(`  _ "github.com/go-sql-driver/mysql"`) }
  if use_json { add(`  "encoding/json"`) }
  if use_fmt { add(`  "fmt"`) }

  add(`)`)

  if use_db == false {
    add(`func escape_print(j []byte)[]byte {`)
    add(` return j`)
    add(`}`)
  }

  add(`func main() {`)
}

func after(lines []string) {
  if use_db == true {

  } else {
    for _, line := range lines {
      if s.Contains(line, "=") {
        value := s.Split(line, "= ")[1]
        if s.Contains(line, "/") {
          r := s.Split(value, "/")
          x := r[0]+`["`+r[1]+`"]`
          add(`respond, _ := json.Marshal(`+x+`)`)
        } else {
          add(`respond, _ := json.Marshal(main["`+value+`"])`)
        }
        add(`respond = escape_print(respond)`)
        add(`fmt.Println(string(respond))`)
      }
    }

  }
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
  init_struct(lines)
  after(lines)
  write_file()
}
