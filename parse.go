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

func build_structs() {

}

func assign_values() {

}

func parse() {
  contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  var lines = strings.Split(string(contents), "\n")
  for _,line := range lines {
    if(strings.Contains(line, " ")) {
      symbols := []byte(line)
      deep := 0
      for _,symbol := range symbols {
        if(symbol == 32) {
          deep++
        }
      }
      println(deep/2)
    }
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
