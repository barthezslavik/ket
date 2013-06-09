package main

import (
  "github.com/hoisie/web"
  "fmt"
  //  "io/ioutil"
  "strings"
  "os"
  "os/exec"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func output(file_name string) string {
  pwd, err := os.Getwd()
  check(err)

  if(file_name != "favicon.ico") {
    file_name = strings.Replace(file_name, ".ket", "", -1)
    var command = pwd + "/" + file_name
    fmt.Println(command)
    dateCmd := exec.Command(string(command))
    dateOut, err := dateCmd.Output()
    check(err)
    result := fmt.Sprint("<pre>", string(dateOut), "</pre>")
    return result
  }
  return ""
}

func main() {
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
