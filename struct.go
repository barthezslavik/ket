package main
import (
  "fmt"
  "encoding/json"
)

func main() {
user := map[string]interface{}{}

second := map[string]interface{}{}

hello := map[string]interface{}{}
user["second"] = second
second["name"] = "slavik"
hello["super"] = "hero"
j, _ := json.Marshal(user)
fmt.Println(string(j))
}
