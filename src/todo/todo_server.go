package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type TodoItem struct {
  Title string `json:"title"`
  Order string `json:"order"`
}

type Todos struct {
  Items []TodoItem
}

func (t Todos) Add(item TodoItem) {
  t.Items = append(t.Items, item)
}

var todos Todos = Todos{}

type TodoServer struct {
}

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  if (r.Method == "POST") {
    var todoItem TodoItem
    body, _ := ioutil.ReadAll(r.Body)

    json.Unmarshal(body, &todoItem)
    todos.Add(todoItem)

    fmt.Fprintf(w, string(body))

  } else if (r.Method == "DELETE") {
    todos.Items = make([]TodoItem, 0)

  } else if (r.Method == "GET") {
    resultTodos, _ := json.Marshal(todos.Items)
    fmt.Fprintf(w, string(resultTodos))
  }
}

