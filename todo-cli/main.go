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
var scanner = bufio.NewScanner(os.Stdin)

func AddTodo(text string, done bool) {
	id := nextId
	nextId++
	todo := Todo{
		ID:   id,
		Text: text,
		Done: done,
	}
	keep = append(keep, todo)
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

func ListTodo() []Todo {
	return keep
}

// A slice expression is [start:end]
// [:i] = give me everything from beginning up to i, but don't include i
// [i:] give me everything starting at index i.
func DeleteTodo(id int) int {
	for i, v := range keep {
		if v.ID == id {
			keep = append(keep[:i], keep[i+1:]...)
			return id
		}
	}
	return id
}

func readString(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}

func readInt(scanner *bufio.Scanner) (int, error) {
	scanner.Scan()
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return num, err
}

func readBool(scanner *bufio.Scanner) (bool, error) {
	scanner.Scan()
	bc, err := strconv.ParseBool(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	return bc, err
}

func main() {
	for {
		fmt.Println(
			"1. Add Todo",
			"\n2. List Todo",
			"\n3. Delete Todo",
			"\n4. Update Todo",
			"\n5. Exit\n",

			"\nChoose an option",
		)
		op1, _ := readInt(scanner)

		switch op1 {
		case 1:
			fmt.Println("enter todo:")
			text := readString(scanner)

			fmt.Println("enter todo status, true and false only:")
			bc, _ := readBool(scanner)
			AddTodo(text, bc)

		case 2:
			todos := ListTodo()
			for _, v := range todos {
				fmt.Println("\nID:", v.ID, "\nTodo:", v.Text, "\nStatus:", v.Done)
			}

		case 3:
			fmt.Println("enter id number to delete")
			id, _ := readInt(scanner)
			DeleteTodo(id)

		case 4:
			fmt.Println("enter the todo number to update")
			id, _ := readInt(scanner)

			fmt.Println("enter Update todo text")
			text := readString(scanner)

			fmt.Println("enter update todo status")
			uc, _ := readBool(scanner)

			UpdateTodo(id, text, uc)

		case 5:
			os.Exit(0)
		}
	}

}
