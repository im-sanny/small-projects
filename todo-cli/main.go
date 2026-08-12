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
var nextId = 3

func AddTodo(text string, done bool) (string, bool) {
	id := nextId
	nextId++
	todo := Todo{
		ID:   id,
		Text: text,
		Done: done,
	}
	keep = append(keep, todo)

	return todo.Text, todo.Done
}

func UpdateTodo(id int, todo string, done bool) {
	for i, v := range keep {
		if v.ID == id {
			keep[i].Text = todo
			keep[i].Done = done
			return
		}
	}
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
	var op1 int
	for {
		fmt.Println(
			"1. Add Todo",
			"\n2. List Todo",
			"\n3. Delete Todo",
			"\n4. Update Todo",
			"\n5. Exit\n",

			"\nChoose an option",
		)
		fmt.Scanln(&op1)

		switch op1 {
		case 1:

			fmt.Println("enter todo:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			text := scanner.Text()

			fmt.Println("enter todo status, true and false only:")
			bscan := bufio.NewScanner(os.Stdin)
			bscan.Scan()
			bconv := bscan.Text()
			bc, _ := strconv.ParseBool(bconv)
			AddTodo(text, bc)
		case 2:
			ListTodo()
		case 3:
			DeleteTodo()
		case 4:
			var uId int
			fmt.Println("enter the todo number to update")
			fmt.Scanln(&uId)
			fmt.Println("enter Update todo text")
			uscan := bufio.NewScanner(os.Stdin)
			uscan.Scan()
			utext := uscan.Text()

			fmt.Println("enter update todo status")
			udscan := bufio.NewScanner(os.Stdin)
			udscan.Scan()
			udconv := udscan.Text()
			uc, _ := strconv.ParseBool(udconv)
			UpdateTodo(uId, utext, uc)
		case 5:
			os.Exit(0)
		}
	}

}
