package main

import (
  "fmt"
  "strconv"
)

type TodoItem struct {
  Title string `json:"title"`
  Order string `json:"order"`
  Completed bool `json:"completed"`
  Url string `json:"url"`
  Id string `json:"id"`
}

func NewTodoItem(id string) TodoItem {
  url := fmt.Sprintf("http://localhost:8080/" + id)
  return TodoItem{Completed: false, Id: id, Url: url}
}

type Todos struct {
  Items []TodoItem
}

func (t *Todos) Add(item TodoItem) {
  t.Items = append(t.Items, item)
}

func (t *Todos) DeleteAll() {
  t.Items = make([]TodoItem, 0)
}

type IdGenerator struct {
  id int
}

func (g *IdGenerator) Generate() string {
  current_id := g.id
  g.id += 1
  return strconv.Itoa(current_id)
}
