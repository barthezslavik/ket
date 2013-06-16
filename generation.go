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

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func add(new_line string)[]string {
  content = append(content, new_line)
  return content
}

func init_struct(lines []string) {
  for _, line := range lines {
    if s.Contains(line, ":") { continue }
    if len(line) == 0 { continue }
    line = s.Replace(line, " ", "", -1)
    add(line+` := map[string]interface{}{}`)
  }
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

func build_parent(lines []string, line string, index int, z int) {
  current_indent := check_indent(lines[index])
  parent_indent := check_indent(lines[index-z])
  parent := lines[index-z]
  parent = clear(parent)
  line = clear(line)
  key_value := s.Split(line, ":")

  if current_indent > parent_indent {
    add(parent+`["`+key_value[0]+`"] = "`+key_value[1]+`"`)
  } else {
    build_parent(lines, line, index, z+1)
  }
}

func build_parent2(lines []string, line string, index int, z int) {
  current_indent := check_indent(lines[index])
  if current_indent == 0 { return }
  parent_indent := check_indent(lines[index-z])
  parent := lines[index-z]
  parent = clear(parent)
  line = clear(line)
  if current_indent > parent_indent {
    add(parent+`["`+line+`"] = `+line)
  } else {
    build_parent2(lines, line, index, z+1)
  }
}

func build(current_line string, lines []string) {
  for index, line := range lines {
    if current_line != line { continue }
    build_parent(lines, line, index, 1)
  }
}

func build2(current_line string, lines []string) {
  for index, line := range lines {
    if current_line != line { continue }
    build_parent2(lines, line, index, 1)
  }
}

func init_values(lines []string) {
  for _, line := range lines {
    if !s.Contains(line, ":") { continue }
    build(line, lines)
  }
}

func init_relations(lines []string) {
  for _, line := range lines {
    if s.Contains(line, ":") { continue }
    build2(line, lines)
  }
}

func before() {
  add(`package main`)
  add(`import (`)
  add(`  "fmt"`)
  add(`  "encoding/json"`)
  add(`)`)
  add(`func escape_print(j []byte)[]byte {`)
  add(` return j`)
  add(`}`)
  add(`func main() {`)
}

func after() {
  add(`j, _ := json.Marshal(user)`)
  add(`j = escape_print(j)`)
  add(`fmt.Println(string(j))`)
  add(`}`)
}

func write_file() {
  f, _ := os.Create(file+".go")
  for _, line := range content {
    f.WriteString(line+"\n")
  }
}

func main() {
  before()
  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  init_struct(lines)
  init_values(lines)
  init_relations(lines)
  after()
  write_file()
}
