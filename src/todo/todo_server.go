package main

import (
  "fmt"
  "net/http"
)

type TodoServer struct {
}

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello world! %s", r.URL.Path[1:])
}

