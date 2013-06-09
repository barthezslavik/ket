package main
import (
  "os"
  "io/ioutil"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {
  data := `package main
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
}`
  content := []byte(data)
  err := ioutil.WriteFile("/tmp/"+os.Args[1]+".go", content, 0644)
  check(err)
}
