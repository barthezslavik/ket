package main

import "fmt"
import "os/exec"

func main() {
  dateCmd := exec.Command("date")
  dateOut, err := dateCmd.Output()
  if err != nil {
    panic(err)
  }
  fmt.Println(string(dateOut))
}
