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
second := map[string]interface{}{}
someone := map[string]interface{}{}
hello := map[string]interface{}{}
user["first"] = first
first["second"] = second
second["name"] = "slavik"
second["age"] = "21"
user["someone"] = someone
someone["hello"] = hello
hello["super"] = "hero"
j, _ := json.Marshal(user)
j = escape_print(j)
fmt.Println(string(j))
}
