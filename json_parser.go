package main
import (
  //"os"
  //"io/ioutil"
  "fmt"
  //"encoding/json"
  //"strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  //contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  //var lines = strings.Split(string(contents), "\n")

  a := make([]string, 1)
  a = append(a, `second  := map[string]interface{} { "name": "slavik", "age": 21 }`)
  a = append(a, `hello   := map[string]interface{} { "super": "hero"}`)
  a = append(a, `someone := map[string]interface{} { "hello": hello }`)
  a = append(a, `first   := map[string]interface{} { "second": second }`)
  a = append(a, `user    := map[string]interface{} { "first": first, "someone": someone }`)

  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
     a[i], a[j] = a[j], a[i]
  }

  //j, _ := json.Marshal(user)

  fmt.Println(a)
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
