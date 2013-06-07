package main

import "fmt"

func main() {
  var first, second int

  user := make([][]string, first)
  for i := range user {
    user[i] = make([]string, second)
  }

  fmt.Println(user[0][0])
}

//user := make(map[string][string]string)
//user["first"]["name"] = "slavik"
//fmt.Println(user["first"]["name"])
