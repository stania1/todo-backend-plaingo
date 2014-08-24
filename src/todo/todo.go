package main

type TodoItem struct {
  Title string `json:"title"`
  Order string `json:"order"`
  Completed bool `json:"completed"`
  Url string `json:"url"`
}

func NewTodoItem() TodoItem {
  return TodoItem{Completed: false}
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
