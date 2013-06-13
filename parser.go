package main
import (
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func parse_object(a []string) []string {
  a = append(a, `
  user    := map[string]interface{} { "first": first, "someone": someone }`)
  a = append(a, `
  someone := map[string]interface{} { "hello": hello }`)
  a = append(a, `
  first   := map[string]interface{} { "second": second }`)
  a = append(a, `
  second  := map[string]interface{} { "name": "slavik", "age": 21 }`)
  a = append(a, `
  hello   := map[string]interface{} { "super": "hero"}`)
  a = append(a, `
  println(user)`)
  return a
}

func main() {
  a := make([]string, 0)
  a = append(a, `
  package main`)
  a = append(a, `
  func main() {`)
  a = parse_object(a)
  a = append(a, `
  }`)

  //for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
  //  a[i], a[j] = a[j], a[i]
  //}

  f, err := os.Create(os.Args[1]+".go")
  check(err)
  for i := 0; i < len(a); i++ {
    f.WriteString(a[i])
  }

  //j, _ := json.Marshal(user)
  //fmt.Println(a)

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
