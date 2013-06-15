package main
import (
  "fmt"
  "encoding/json"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func dont_print(j []byte)[]byte {
  return j
}

func main() {
  user := map[string]interface{}{}
  first := map[string]interface{}{}
  someone := map[string]interface{}{}
  second := map[string]interface{}{}
  hello := map[string]interface{}{}

  second["name"] = "slavik"
  second["age"] = 21
  hello["super"] = "hero"

  someone["hello"] = hello
  first["second"] = second
  user["first"] = first
  user["someone"] = someone

  j, err := json.Marshal(user)
  check(err)

  j = dont_print(j)
  fmt.Println(string(j))
}
