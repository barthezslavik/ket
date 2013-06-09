/*

user
  first
    second
      name:slavik
      age:26

= user/first/second/name

*/

package main
import "fmt"

type user struct {
  _name string
  first
}

type first struct {
  _name string
  second
}

type second struct {
  name string
  age int
}

func main() {
  data :=
    user {
      "first", first {
        "second", second {
          name: "slavik"}}}

  fmt.Println(data)
}