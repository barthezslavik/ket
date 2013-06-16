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

  build := exec.Command(pwd + "/build", file_name)
  build_out, err := build.Output()
  check(err)
  println(build_out)

  run := exec.Command(pwd + "/run", file_name)
  run_out, err := run.Output()
  check(err)
  result := fmt.Sprint(string(run_out))
  return result
}

func main() {
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
