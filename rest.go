package main

import (
  "github.com/emicklei/go-restful"
  "net/http"
)

func main() {
  ws := new(restful.WebService)
  ws.Route(ws.GET("/hello").To(hello))
  restful.Add(ws)
  http.ListenAndServe(":9999", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
  resp.Write([]byte("world"))
}
