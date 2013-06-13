package main
import (
  //"os"
  //"io/ioutil"
  "fmt"
  "encoding/json"
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

second := map[string]interface{} { "name": "slavik", "age": 21 }
first := map[string]interface{} { "second":second }
hello := map[string]interface{} { "super": "hero"}
someone := map[string]interface{} { "hello":hello }
user := map[string]interface{} { "first":first, "someone":someone }

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
