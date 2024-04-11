package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"todo/assets/structs"
)

// Funktion, um das heutige Datum im gewünschten Format zu erhalten
func todayDate() string {
	now := time.Now().Local()
	date := now.Format("2006-01-02")
	return date
}

// Funktion zum Hinzufügen einer neuen Aufgabe zur JSON-Datei
func AddEntry(todo string) {
	// Neue Aufgabe erstellen
	newEntry := structs.TodoList{
		Created: todayDate(),
		Done:    "",
		Desc:    todo,
	}

	// JSON-Datei öffnen (oder erstellen, falls nicht vorhanden) im Schreibmodus, ohne Anhängen
	file, err := os.OpenFile("todo_list.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	// JSON-Daten aus der Datei lesen
	var tasks []structs.TodoList
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err.Error() != "EOF" {
		fmt.Println("Fehler beim Lesen der JSON-Daten:", err)
		return
	}

	// Neue Aufgabe zur Liste hinzufügen
	tasks = append(tasks, newEntry)

	// Zurück zum Anfang der Datei springen und Inhalt löschen
	if err := file.Truncate(0); err != nil {
		fmt.Println("Fehler beim Zurücksetzen des Speichers:", err)
		return
	}

	// Zurück zum Anfang der Datei springen, um die neuen Daten zu schreiben
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("Fehler beim Zurücksetzen des Dateizeigers:", err)
		return
	}

	// Neue Aufgabenliste in die Datei schreiben
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Fehler beim Schreiben in die Datei:", err)
		return
	}

	fmt.Println("Eine weitere Aufgabe wurde hinzugefügt.")
}
