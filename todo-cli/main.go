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

func AddTodo(a string) {
	fmt.Println(a)
}

func main() {
	AddTodo("learn go")
}
