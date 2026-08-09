package main

import "fmt"

type Todo struct {
	ID   int
	Text string
	Done bool
}

var keep = []Todo{
	{ID: 1, Text: "write 10 lines of code", Done: false},
	{ID: 2, Text: "don't waste much time in social", Done: false},
}

func AddTodo(text string, done bool) (string, bool) {
	id := len(keep) + 1
	todo := Todo{
		ID:   id,
		Text: text,
		Done: done,
	}
	keep = append(keep, todo)

	return todo.Text, todo.Done
}

func ListTodo() {
	for _, v := range keep {
		fmt.Println(v.ID, v.Text)
	}
}

func DeleteTodo(id int) {
	for i, v := range keep {
		if v.ID == id {
			keep = append(keep[:i], keep[i+1:]...)
			return
		}
	}
}

func main() {
	AddTodo("learn go", false)
	DeleteTodo(1)
	ListTodo()
}
