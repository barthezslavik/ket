package main
import (
  "os"
  //"fmt"
  //"encoding/json"
  //"strings"
  //"bufio"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  //contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  //var lines = strings.Split(string(contents), "\n")

  a := make([]string, 0)
  a = append(a, `second  := map[string]interface{} { "name": "slavik", "age": 21 }`)
  a = append(a, `hello   := map[string]interface{} { "super": "hero"}`)
  a = append(a, `someone := map[string]interface{} { "hello": hello }`)
  a = append(a, `first   := map[string]interface{} { "second": second }`)
  a = append(a, `user    := map[string]interface{} { "first": first, "someone": someone }`)

  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
  }

  //for i := 0; i < len(a); i++ {
  //  fmt.Println(a[i])
    //err := ioutil.WriteFile("/tmp/www", string(a[i]), 0644)
    //check(err)
  //}
  //j, _ := json.Marshal(user)

  //fmt.Println(a)

  f, err := os.Create("/tmp/dat2")
  check(err)
  f.WriteString("writes\n")
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
