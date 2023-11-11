package main

import (
	"fmt"
	"os"
	todo "todo_cli_app"
)

const (
	// hidden file
	todoFile = ".todos.json"
)

func main() {
	fmt.Println("hello, world")

	// 1. Read and write todos into json file
	// 2. basic operations for todo are
	// 2. a. to list all the todos
	// 2. b. to add a todo to the list
	// 2. c. delete
	// 2. d. Complete todo check
	// 2. e. invalid commands
	// 3. to read/ scann the the inputss
	// 4. Print the todos

	todos := &todo.Todos{}
	// if error in loading, exit the prrogram
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

}
