package utils

import (
	"encoding/json"
	"fmt"
	"time"
	"todo/assets/filethings"
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
	tasks, file := filethings.OpenAndReadFile()
	if file != nil {
		defer file.Close()
	} else {
		return
	}

	// Neue Aufgabe erstellen
	newEntry := structs.TodoList{
		Created: todayDate(),
		Done:    "",
		Desc:    todo,
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
		fmt.Println("Fehler beim Zurücksetzen des Speicherzeigers:", err)
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
