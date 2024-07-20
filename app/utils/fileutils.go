package utils

import (
	"encoding/json"
	"os"
	"github.com/yigitkalan/gotodo/app/models"
)

const DEFAULT_FILENAME = "tasks.json"

// Saves the tasks to the file specified in the TODO_PATH or default path in json format.
func SaveToFile(todos []models.Item) error {
	appPath := verifyPaths()
	filePath := appPath + "/" + DEFAULT_FILENAME
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(todos)
}

// Verifies if the default path exists, if not creates it and returns it.
//
// To change the default path set the environment variable TODO_PATH to the desired path.
func verifyPaths() string {
	path := os.Getenv("TODO_PATH")
	home, _ := os.UserHomeDir()
	if path == "" {
		path = home + "/.local/share/todo"
	}
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
	return path

}

// Loads the tasks from the file specified in the TODO_PATH or default path in json format.
func LoadFromFile() ([]models.Item, error) {
	appPath := verifyPaths()
	filePath := appPath + "/" + DEFAULT_FILENAME
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var tasks []models.Item
	err = decoder.Decode(&tasks)
	return tasks, err
}
