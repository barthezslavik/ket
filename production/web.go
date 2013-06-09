package main

import (
  "github.com/hoisie/web"
  "fmt"
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

  cmd := exec.Command(pwd + "/" + file_name)
  out, err := cmd.Output()
  check(err)
  result := fmt.Sprint("<pre>", string(out), "</pre>")
  return result
}

func main() {
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
