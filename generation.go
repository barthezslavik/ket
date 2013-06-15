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

func main() {
  second  := map[string]interface{}{}
  second["name"] = "slavik"
  second["age"] = 21
  hello   := map[string]interface{}{}
  hello["super"] = "hero"
  someone := map[string]interface{}{}
  someone["hello"] = hello
  first   := map[string]interface{}{}
  first["second"] = second
  user    := map[string]interface{}{}
  user["first"] = first
  user["someone"] = someone
  user["first"] = second
  user["someone"] = someone

  j, _ := json.Marshal(user)
  fmt.Println(string(j))

}

/*
user
  first
    second
      name:slavik
      age:21
  someone
    hello
      super:hero
*/
