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
  Items map[string]TodoItem
}

func (t *Todos) Add(item TodoItem) {

  if (len(t.Items) == 0) {
    t.Items = make(map[string]TodoItem)
  }

  t.Items[item.Id] = item
}

func (t *Todos) DeleteAll() {
  t.Items = make(map[string]TodoItem)
}

func (t Todos) Get(id string) TodoItem {
  return t.Items[id]
}

func (t Todos) AsArray() []TodoItem {
  m := t.Items
  v := make([]TodoItem, 0, len(m))

  for  _, value := range m {
    v = append(v, value)
  }
  return v
}

type IdGenerator struct {
  id int
}

func (g *IdGenerator) Generate() string {
  current_id := g.id
  g.id += 1
  return strconv.Itoa(current_id)
}
