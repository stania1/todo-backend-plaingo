package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type TodoServer struct {
  todos Todos
}

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  if (r.Method == "POST") {
    var todoItem TodoItem
    body, _ := ioutil.ReadAll(r.Body)

    json.Unmarshal(body, &todoItem)
    server.todos.Add(todoItem)

    fmt.Fprintf(w, string(body))

  } else if (r.Method == "DELETE") {
    server.todos.DeleteAll()

  } else if (r.Method == "GET") {
    resultTodos, _ := json.Marshal(server.todos.Items)
    fmt.Fprintf(w, string(resultTodos))
  }
}

