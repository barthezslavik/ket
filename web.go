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

  file_name = strings.Replace(file_name, ".ket", "", -1)
  dateCmd := exec.Command(pwd + "/run", file_name)
  dateOut, err := dateCmd.Output()
  check(err)
  result := fmt.Sprint("<pre>", string(dateOut), "</pre>")
  return result
}

func root() string {
  return ""
}

func main() {
  web.Get("/", root)
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
