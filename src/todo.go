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

func corsHandler(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, GET, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
    if (r.Method == "OPTIONS") {
      // preflight request. just return 200 OK.
    } else {
      h.ServeHTTP(w,r)
    }
  }
}

func main() {
  server := new(TodoServer)
  fmt.Println("test")
  http.HandleFunc("/", corsHandler(server))
  http.ListenAndServe(":8080", nil)
}
