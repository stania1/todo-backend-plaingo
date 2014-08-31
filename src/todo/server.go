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
    server.HandlePOST(w, r)

  } else if (r.Method == "DELETE") {
    server.HandleDELETE(w, r)

  } else if (r.Method == "GET" && r.URL.Path == "/") {

    server.HandleGETRoot(w, r)

  } else if (r.Method == "GET" && r.URL.Path != "/") {
    server.HandleGETSingle(w, r)

  } else if (r.Method == "PATCH") {
    server.HandlePATCH(w, r)
  }
}

func (server *TodoServer) HandlePOST(w http.ResponseWriter, r *http.Request) {
  id := server.idGenerator.Generate()
  var todoItem TodoItem = NewTodoItem(id)
  body, _ := ioutil.ReadAll(r.Body)

  json.Unmarshal(body, &todoItem)
  server.todos.Add(todoItem)

  fmt.Fprintf(w, todoItem.String())
}

func (server *TodoServer) HandlePATCH(w http.ResponseWriter, r *http.Request) {
  todoId := r.URL.Path[1:]
  todoItem := server.todos.Get(todoId)

  body, _ := ioutil.ReadAll(r.Body)
  json.Unmarshal(body, &todoItem)

  server.todos.Update(todoItem)

  fmt.Fprintf(w, todoItem.String())
}

func (server *TodoServer) HandleDELETE(w http.ResponseWriter, r *http.Request) {
  server.todos.DeleteAll()
}

func (server *TodoServer) HandleGETRoot(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, server.todos.String())
}

func (server *TodoServer) HandleGETSingle(w http.ResponseWriter, r *http.Request) {
  todoId := r.URL.Path[1:]
  todoItem := server.todos.Get(todoId)

  fmt.Fprintf(w, todoItem.String())
}
