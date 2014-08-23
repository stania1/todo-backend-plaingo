package main

type TodoItem struct {
  Title string `json:"title"`
  Order string `json:"order"`
}

type Todos struct {
  Items []TodoItem
}

func (t *Todos) Add(item TodoItem) {
  t.Items = append(t.Items, item)
}


