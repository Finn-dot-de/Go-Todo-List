package filethings

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/assets/structs"
)

// Funktion zum Öffnen und Lesen der JSON-Datei
func OpenAndReadFile() ([]structs.TodoList, *os.File) {
	// JSON-Datei öffnen (oder erstellen, falls nicht vorhanden) im Schreibmodus, ohne Anhängen
	file, err := os.OpenFile("todo_list.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return nil, nil
	}

	// JSON-Daten aus der Datei lesen
	var tasks []structs.TodoList
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		fmt.Println("Fehler beim Lesen der JSON-Daten:", err)
		return nil, nil
	}

	return tasks, file
}
