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
  if current_indent == 0 { return }
  parent_indent := check_indent(lines[index-z])
  parent := lines[index-z]
  parent = clear(parent)
  line = clear(line)

  if current_indent > parent_indent {
    if f == "n" { add(parent+`["`+line+`"] = `+line) }
    if f == "k" { add(parent+`["`+s.Split(line, ":")[0]+`"] = "`+s.Split(line, ":")[1]+`"`) }
  } else {
    build(lines, line, index, z+1, f)
  }
}

func init_struct() {
  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  for index, line := range lines {
    if s.Contains(line, ":") {
      build(lines, line, index, 1, "k")
    } else {
      if len(line)>0 { add(clear(line)+` := map[string]interface{}{}`) }
      build(lines, line, index, 1, "n")
      if index == 0 { struct_name = line }
    }
  }
}

func before() {
  add(`package main`)
  add(`import (`)
  if len(struct_name)>0 {
    add(`  "fmt"`)
    add(`  "encoding/json"`)
  }
  add(`)`)
  add(`func escape_print(j []byte)[]byte {`)
  add(` return j`)
  add(`}`)
  add(`func main() {`)
}

func after() {
  if len(struct_name)>0 {
    add(`j, _ := json.Marshal(`+struct_name+`)`)
    add(`j = escape_print(j)`)
    add(`fmt.Println(string(j))`)
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
