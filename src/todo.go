package main

import (
  "fmt"
  "net/http"
)

type MyServer struct {
}

func (server *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  fmt.Println(r.Method)
  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
  }

  if r.Method == "OPTIONS" {
    return
  }

  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, GET, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
  fmt.Fprintf(w, "Hello world! %s", r.URL.Path[1:])
}

func corsHandler(h http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    if (r.Method == "OPTIONS") {
      //handle preflight in here
    } else {
      h.ServeHTTP(w,r)
    }
  }
}

func main() {
  server := new(MyServer)
  fmt.Println("test")
  http.HandleFunc("/", corsHandler(server))
  http.ListenAndServe(":8080", nil)
}
