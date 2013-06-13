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

  //first := //

  var a [5]string
  a[0] = `second  := map[string]interface{} { "name": "slavik", "age": 21 }`
  a[1] = `hello   := map[string]interface{} { "super": "hero"}`
  a[2] = `someone := map[string]interface{} { "hello": hello }`
  a[3] = `first   := map[string]interface{} { "second": second }`
  a[4] = `user    := map[string]interface{} { "first": first, "someone": someone }`

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
