package utils

import (
	"bufio"
	"fmt"
	m "github.com/yigitkalan/gotodo/app/models"
	"os"
	"strconv"
)

var tasks []m.Item

func Operate(operation m.Operation) {
	switch operation {
	case m.Add:
		AddTask()
	case m.Complete:
		CompleteTask()
	case m.Exit:
		Exit()
	case m.Help:
		ShowHelp()
	case m.List:
		ListTasks()
	case m.Remove:
		RemoveTask()
	default:
		panic(fmt.Sprintf("unexpected main.Operation: %#v", operation))
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
		fmt.Println(i, "-", m.Operation(i))
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
	tasks = append(tasks, m.Item{ID: len(tasks), Description: description, Completed: false})

}

func Exit() {
	err := SaveToFile(tasks)
	if err != nil {
		fmt.Println("Error saving tasks to file")
	}else {
		os.Exit(0)
	}
}

func SetTasks(newTasks []m.Item) {
    tasks = newTasks
}
