package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Operation int8

const (
	Add Operation = iota
	Remove
	Complete
	List
	Exit
	Help
)

func main() {
	fmt.Println("Welcome to the task manager!")
	ShowHelp()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Select an option: ")
	for scanner.Scan() {
		line := scanner.Text()
		option, err := strconv.Atoi(line)
		if err != nil || option < 0 || option > int(Help) {
			fmt.Println("Invalid option, please try again")
			fmt.Print("Select an option: ")
			continue
		}

		fmt.Print("Select an option: ")
	}
}

func (operation Operation) Operate() {
	switch operation {
	case Add:
		AddTask()
	case Complete:
		CompleteTask()
	case Exit:
		os.Exit(0)
	case Help:
		ShowHelp()
	case List:
		ListTasks()
	case Remove:
		RemoveTask()
	default:
		panic(fmt.Sprintf("unexpected main.Operation: %#v", operation))
	}
}

func RemoveTask() {
	panic("unimplemented")
}

func ListTasks() {
	panic("unimplemented")
}

func ShowHelp() {
	fmt.Println("\nThe following commands are available:")
	fmt.Println("0 - Add a task")
	fmt.Println("1 - Remove a task")
	fmt.Println("2 - Complete a task")
	fmt.Println("3 - List all tasks")
	fmt.Println("4 - Exit")
	fmt.Println("5 - Help")
}

func CompleteTask() {
	panic("unimplemented")
}

func AddTask() {
	panic("unimplemented")
}
