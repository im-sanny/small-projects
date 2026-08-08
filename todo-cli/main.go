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

func AddTodo(a int, b string, c bool) (int, string, bool) {

	id := len(keep) + 1
	Todo.ID[] = id
	keep = append(keep, Todo{})

	return keep
}

func main() {
	AddTodo(len(keep)+1, "learn go", true)
	fmt.Println(len(keep))
}
