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

func main() {
  contents,_ := ioutil.ReadFile(os.Args[1]+".ket")
  var lines = strings.Split(string(contents), "\n")
  println(lines[0])
}
