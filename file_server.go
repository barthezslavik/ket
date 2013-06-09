package main

import (
  "io"
  "net/http"
  "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
  //file_name = strings.Replace(file_name, ".ket", ".go", -1)
  //var command = pwd + "/build"// -v="// + file_name
  //dateCmd := exec.Command(string(command))
  //dateOut, err := dateCmd.Output()
  //check(err)
  //result := fmt.Sprint("<pre>", string(dateOut), "</pre>")
  io.WriteString(w, "result")
}

func main() {
  http.Handle("/", http.FileServer(http.Dir("./")))
  http.HandleFunc("*", HelloServer)
  err := http.ListenAndServe(":9999", nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
  }
}
