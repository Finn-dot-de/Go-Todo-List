package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/assets/structs"
)

// ReadJsonFile liest die Aufgabenliste aus einer JSON-Datei und gibt sie zurück.
// Es gibt auch einen Fehler zurück, falls beim Lesen der Datei oder beim Decodieren der Daten ein Problem auftritt.
func ReadJsonFile() ([]structs.TodoList, error) {
	// Öffnen der JSON-Datei im Lesemodus
	file, err := os.OpenFile("todo_list.json", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		// Fehlerbehandlung, falls das Öffnen der Datei fehlschlägt
		return nil, fmt.Errorf("fehler beim öffnen der datei: %s", err)
	}
	defer file.Close()

	// JSON-Dekoder für die Datei erstellen
	dekoder := json.NewDecoder(file)

	// Aufgabenliste erstellen, um die gelesenen Daten zu speichern
	var aufgaben []structs.TodoList
	// JSON-Daten in die Aufgabenliste decodieren
	if err := dekoder.Decode(&aufgaben); err != nil {
		// Fehlerbehandlung, falls das Decodieren der Daten fehlschlägt
		return nil, fmt.Errorf("fehler beim lesen der JSON-daten: %s", err)
	}

	// Rückgabe der gelesenen Aufgabenliste
	return aufgaben, nil
}

// PrintTasks gibt die übergebene Aufgabenliste formatiert auf der Konsole aus.
func PrintTasks(tasks []structs.TodoList) {
	// Iteration über die Aufgabenliste
	for i, aufgabe := range tasks {
		// Status der Aufgabe festlegen
		Check := "[ ]"
		if aufgabe.Done != "" {
			Check = "[X]"
		}
		// Aufgabe formatiert ausgeben
		fmt.Printf("| %d | %3s | %10s | %10s | %8s |\n", i+1, Check, aufgabe.Created, aufgabe.Done, aufgabe.Desc)
	}
}
