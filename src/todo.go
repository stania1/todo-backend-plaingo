package main

import (
  "fmt"
  "net/http"
)

func handler(writer http.ResponseWriter, reader *http.Request) {
  fmt.Fprintf(writer, "Hello world! %s", reader.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
