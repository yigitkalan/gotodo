package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/yigitkalan/gotodo/app/models"
	"github.com/yigitkalan/gotodo/app/utils"
)


func main() {
	fmt.Println("Welcome to the task manager!")
	operations.ShowHelp()
	AppLoop()
}

func AppLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("--> ")
	for scanner.Scan() {
		line := scanner.Text()
		option, err := strconv.Atoi(line)
		if err != nil || option < 0 || option > int(models.Help) {
			fmt.Println("Invalid option, please try again")
			fmt.Print("--> ")
			continue
		}
		operation := models.Operation(option)

		operations.Operate(operation)

		fmt.Println()
		fmt.Print("--> ")
	}

}

