package main

import (
  "bytes"
  "fmt"
  "log"
  "os/exec"
  "strings"
)

func main() {
  cmd := exec.Command("go build /home/slavik/go/ket/struct.go")
  cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer
  cmd.Stdout = &out
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("in all caps: %q\n", out.String())
}
