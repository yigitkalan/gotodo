package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Operation int8

const (
	Exit Operation = iota
	List
	Add
	Complete
	Remove
	Help
)
var operationNames = map[Operation]string{
    Exit:     "Exit",
    List:     "List",
    Add:      "Add",
    Complete: "Complete",
    Remove:   "Remove",
    Help:     "Help",
}

type Task struct {
	ID          int
	Description string
	Completed   bool
}

var tasks []Task

func (o Operation) String() string {
    return operationNames[o]
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

func main() {
	fmt.Println("Welcome to the task manager!")
	ShowHelp()
	AppLoop()
}

func AppLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("--> ")
	for scanner.Scan() {
		line := scanner.Text()
		option, err := strconv.Atoi(line)
		if err != nil || option < 0 || option > int(Help) {
			fmt.Println("Invalid option, please try again")
			fmt.Print("--> ")
			continue
		}
		operation := Operation(option)

		operation.Operate()

		fmt.Println()
		fmt.Print("--> ")
	}

}


func RemoveTask() {
	idToRemove := GetIdFromUser()
	removeTaskById(idToRemove)
}

func removeTaskById(idToRemove int) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to remove")
		fmt.Println()
		return
	}

	for i, task := range tasks {
		if task.ID == idToRemove {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
		if i == len(tasks)-1 {
			fmt.Println("Task not found")
		}
	}
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Description: %s, Completed: %t\n", task.ID, task.Description, task.Completed)
	}
}

func ShowHelp() {
	fmt.Println("\nThe following commands are available:")
    for i := 0; i < 5; i++ { 
        fmt.Println(i, "-", Operation(i))
    }
}

func CompleteTask() {
	id := GetIdFromUser()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			return
		}
	}
}

func GetIdFromUser() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the ID of the task you want to remove: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Invalid ID, please try again")
		return GetIdFromUser()
	}
	return id
}

func AddTask() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter the description of the task: ")
	scanner.Scan()
	description := scanner.Text()
	tasks = append(tasks, Task{ID: len(tasks), Description: description, Completed: false})

}
