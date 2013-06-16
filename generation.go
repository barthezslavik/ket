package main

import (
  "fmt"
  //  "encoding/json"
  "io/ioutil"
  s "strings"
  "os"
)

var p = fmt.Println
//var file = os.Args[1]+"ket"
var file = "struct"
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

func find_parent(current_line string, lines []string) {
  for index, line := range lines {
    if current_line != line { continue }
    current_indent := check_indent(line)
    parent_indent := check_indent(lines[index-1])
    if current_indent - parent_indent != 2 { continue }
    parent := lines[index-1]
    parent = s.Replace(parent, " ", "", -1)
    line = s.Replace(line, " ", "", -1)
    key_value := s.Split(line, ":")
    add(parent+`["`+key_value[0]+`"] = "`+key_value[1]+`"`)
  }
}

func init_relations(lines []string) {
  for _, line := range lines {
    if !s.Contains(line, ":") { continue }
    find_parent(line, lines)
  }
  //add(`someone["hello"] = hello`)
  //add(`first["second"] = second`)
  //add(`user["first"] = first`)
  //add(`user["someone"] = someone`)
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
  add(`fmt.Println(someone)`)
  add(`fmt.Println(first)`)
  add(`}`)
}

func write_file() {
  f, _ := os.Create(file+".go")
  for _, line := range content {
    f.WriteString(line+"\n")
    //p(line)
  }
}

func main() {
  before()
  contents,_ := ioutil.ReadFile(file+".ket")
  lines := s.Split(string(contents), "\n")
  init_struct(lines)
  init_relations(lines)
  after()
  write_file()
  //p(content)
}
