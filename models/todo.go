package models

type Todo struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
	Done bool   `json:"done"`
}

type TodoForm struct {
	Body string `json:"body"`
	Done bool   `json:"done"`
}
