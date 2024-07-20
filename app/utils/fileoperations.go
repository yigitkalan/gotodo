package operations

import (
	"encoding/json"
	"os"

	"github.com/yigitkalan/gotodo/app/models"
)

func saveToFile(todos []models.Task, filename string) error{

    path := os.Getenv("TODO_PATH")
    if(path == ""){
        path = "$HOME/.local/share/todo"
    }
    file, err := os.Create(path + "/" + filename)
    if err != nil {
        return err
    } 
    defer file.Close()
    encoder := json.NewEncoder(file)
    return encoder.Encode(todos)

}
