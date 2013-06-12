package main
import (
  "os"
  "io/ioutil"
  //"fmt"
  "strings"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func node(line string, deep int, lines[]string) string {
  var output string = "type user struct { _name string first }"
  return output
}

func parent() {

}

func none() {

}

func parse() {
  contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  var lines = strings.Split(string(contents), "\n")
  for _,line := range lines {
    symbols := []byte(line)
    deep := 0
    for _,symbol := range symbols {
      if(symbol == 32) {
        deep++
      }
    }
    new_node := node(line, deep, lines)
    buffer := []byte(string(new_node))
    err := ioutil.WriteFile("output.txt", buffer, 0644)
    check(err)
  }

}

func main() {
  parse()
  /*
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
    _name string
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
        */
      }
