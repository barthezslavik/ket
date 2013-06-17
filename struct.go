package main
import (
  "fmt"
  "encoding/json"
)
func escape_print(j []byte)[]byte {
 return j
}
func main() {
user := map[string]interface{}{}
first := map[string]interface{}{}
user["first"] = first
second := map[string]interface{}{}
first["second"] = second
second["name"] = "slavik"
second["age"] = "21"
someone := map[string]interface{}{}
user["someone"] = someone
hello := map[string]interface{}{}
someone["hello"] = hello
hello["super"] = "hero"
j, _ := json.Marshal(user)
j = escape_print(j)
fmt.Println(string(j))
}
