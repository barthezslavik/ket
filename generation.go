package main
import (
  "fmt"
  //  "encoding/json"
  //  "io/ioutil"
  //  s "strings"
  //  "os"
)

var p = fmt.Println
//var file = os.Args[1]+"ket"
var file = "struct.ket"
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

func init_struct() {
  add(`user := map[string]interface{}{}`)
  add(`first := map[string]interface{}{}`)
  add(`someone := map[string]interface{}{}`)
  add(`second := map[string]interface{}{}`)
  add(`hello := map[string]interface{}{}`)
}

func init_values() {
  add(`second["name"] = "slavik"`)
  add(`second["age"] = 21`)
  add(`hello["super"] = "hero"`)
}

func init_relations() {
  add(`someone["hello"] = hello`)
  add(`first["second"] = second`)
  add(`user["first"] = first`)
  add(`user["someone"] = someone`)
}

func before() {
  add(`package main`)
  add(`import (`)
  add(`  "fmt")`)
  add(`func dont_print(j []byte)[]byte {`)
  add(` return j`)
  add(`}`)
}

func after() {
  add(`j, _ := json.Marshal(user)`)
  add(`j = dont_print(j)`)
}

func main() {
  before()
  init_struct()
  init_values()
  init_relations()
  after()
  //contents,_ := ioutil.ReadFile(file)
  //lines := s.Split(string(contents), "\n")
  p(content)
}
