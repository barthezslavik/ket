package main

import (
  "github.com/hoisie/web"
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func output(file_name string) string {
  pwd, err := os.Getwd()
  check(err)

  fmt.Println(pwd)
  dat, err := ioutil.ReadFile(pwd + "/" + file_name)
  check(err)

  return string(dat)
}

func main() {
  web.Get("/(.*)", output)
  web.Run("0.0.0.0:9999")
}
