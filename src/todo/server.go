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

    fmt.Fprintf(w, todoItem.String())

  } else if (r.Method == "DELETE") {
    server.todos.DeleteAll()

  } else if (r.Method == "GET" && r.URL.Path == "/") {

    fmt.Fprintf(w, server.todos.String())

  } else if (r.Method == "GET" && r.URL.Path != "/") {
    todoId := r.URL.Path[1:]
    todoItem := server.todos.Get(todoId)

    fmt.Fprintf(w, todoItem.String())

  } else if (r.Method == "PATCH") {
    todoId := r.URL.Path[1:]
    todoItem := server.todos.Get(todoId)

    body, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(body, &todoItem)

    server.todos.Update(todoItem)

    fmt.Fprintf(w, todoItem.String())
  }
}

