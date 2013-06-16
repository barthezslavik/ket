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
second["name"] = "slavik"
name:slavik["age"] = "21"
second["age"] = "21"
hello["super"] = "hero"
j, _ := json.Marshal(user)
j = escape_print(j)
fmt.Println(string(j))
}
