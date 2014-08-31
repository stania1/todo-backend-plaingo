package main

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

type TodoServer struct {
  todos Todos
  idGenerator IdGenerator
}

func (server *TodoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

  if (r.Method == "POST") {
    id := server.idGenerator.Generate()
    var todoItem TodoItem = NewTodoItem(id)
    body, _ := ioutil.ReadAll(r.Body)

    json.Unmarshal(body, &todoItem)
    server.todos.Add(todoItem)

    resultTodo, _ := json.Marshal(todoItem)
    fmt.Fprintf(w, string(resultTodo))

  } else if (r.Method == "DELETE") {
    server.todos.DeleteAll()

  } else if (r.Method == "GET") {
    todoItems := server.todos.AsArray()
    resultTodos, _ := json.Marshal(todoItems)

    fmt.Fprintf(w, string(resultTodos))
  }
}

