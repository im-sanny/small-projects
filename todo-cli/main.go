package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	fmt.Println("enter todo:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("enter todo status, true and false only:")
	bscan := bufio.NewScanner(os.Stdin)
	bscan.Scan()
	bconv := bscan.Text()
	bc, _ := strconv.ParseBool(bconv)

	id := len(keep) + 1
	todo := Todo{
		ID:   id,
		Text: scanner.Text(),
		Done: bc,
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
