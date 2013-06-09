package main

import (
  "github.com/hoisie/web"
  "fmt"
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

  if(file_name == "favicon.ico"){
    return ""
  }

  file_name = strings.Replace(file_name, ".ket", ".go", -1)
  var command = pwd + "/build"// -v="// + file_name
  dateCmd := exec.Command(string(command))
  dateOut, err := dateCmd.Output()
  check(err)
  result := fmt.Sprint("<pre>", string(dateOut), "</pre>")
  return result
}

func main() {
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
