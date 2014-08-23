package main

import (
  "fmt"
  "net/http"
  // "encoding/json"
  "io/ioutil"
)

type TodoItem struct {
  Title string
  Order string
}

type TodoServer struct {
}

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  if (r.Method == "POST") {
    // var todoItem TodoItem
    body, _ := ioutil.ReadAll(r.Body)
    // // fmt.Println(string(body))
    //
    // json.Unmarshal(body, &todoItem)
    // fmt.Println(todoItem.Title)
    // // fmt.Printf("%+v", todoItem)

    fmt.Fprintf(w, string(body))
  }
  // fmt.Fprintf(w, "Hello world! %s", r.URL.Path[1:])
}

