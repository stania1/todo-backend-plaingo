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
    // convert Items to an array before
    m := server.todos.Items
    v := make([]TodoItem, 0, len(m))

    for  _, value := range m {
      v = append(v, value)
    }

    // resultTodos, _ := json.Marshal(server.todos.Items)
    resultTodos, _ := json.Marshal(v)

    fmt.Fprintf(w, string(resultTodos))
  }
}

