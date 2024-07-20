package main

import (
	"bufio"
	"fmt"
	"github.com/yigitkalan/gotodo/app/models"
	"github.com/yigitkalan/gotodo/app/utils"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to the task manager!")
	utils.ShowHelp()
	AppLoop()
}

func AppLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("--> ")
	tasks, err := utils.LoadFromFile()
	if err != nil {
		fmt.Println("Error loading tasks from file")
	}
	utils.SetTasks(tasks)
	for scanner.Scan() {
		line := scanner.Text()
		option, err := strconv.Atoi(line)
		if err != nil || option < 0 || option > int(models.Help) {
			fmt.Println("Invalid option, please try again")
			fmt.Print("--> ")
			continue
		}
		utils.Operate(models.Operation(option))
		fmt.Println()
		fmt.Print("--> ")
	}
}
