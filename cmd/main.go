package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func main(){
    fmt.Println("Welcome to the task manager!")
    fmt.Println("Enter 'help' to see the list of commands")
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
    }
}

func (operation Operation) Operate(){
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
	panic("unimplemented")
}

func CompleteTask() {
	panic("unimplemented")
}

func AddTask() {
	panic("unimplemented")
}
