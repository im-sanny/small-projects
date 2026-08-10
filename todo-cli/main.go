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

func AddTodo() (string, bool) {
	var text string
	var done bool

	fmt.Println("enter todo:")
	fmt.Scanln(&text)
	fmt.Println("enter todo status, true and false only:")
	n, err := fmt.Scanln(&done)
	fmt.Println("from n", n, "\nform err:", err)

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
		fmt.Println("ID:", v.ID, "\nText:", v.Text, "\nDone:", v.Done)
	}
}

func DeleteTodo() int {
	var id int
	fmt.Println("enter the id of todo u want to delete:")
	fmt.Scanln(&id)
	for i, v := range keep {
		if v.ID == id {
			keep = append(keep[:i], keep[i+1:]...)
			return id
		}
	}
	return id
}

func main() {
	AddTodo()
	DeleteTodo()
	ListTodo()
}
