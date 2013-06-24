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

func init_struct() {
  add(`main := map[string]interface{}{}`)
  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  for index, line := range lines {
    if s.Contains(line, "=") { return }
    if s.Contains(line, "db") { db(); return }
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

func before() {
  add(`package main`)
  add(`import (`)

  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  fmt_import := false
  json_import := false
  for _, line := range lines {
    if s.Contains(line, "=") {
      fmt_import = true
      json_import = true
    }
  }

  if use_db == false {
    //if fmt_import == true { add(`  "fmt"`) }
    //if json_import == true { add(`  "encoding/json"`) }
  }

  add(`)`)

  if use_db == false {
    add(`func escape_print(j []byte)[]byte {`)
    add(` return j`)
    add(`}`)
  }

  add(`func main() {`)
}

func db() {
  use_db = true
}

func after() {
  if use_db == true {

  } else {
    contents,_ := ioutil.ReadFile(file+".ket")
    lines := s.Split(string(contents), "\n")
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
  before()
  init_struct()
  after()
  write_file()
}
