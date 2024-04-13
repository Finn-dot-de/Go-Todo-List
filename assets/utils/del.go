package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"todo/assets/structs"
)

func Delete(index int) {
	// Öffne die JSON-Datei
	file, err := os.OpenFile("todo_list.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Fehler beim Öffnen der Datei: %v\n", err)
		return
	}
	defer file.Close()

	// Dekodiere den Inhalt der Datei in ein Slice
	var slice []structs.TodoList
	err = json.NewDecoder(file).Decode(&slice)
	if err != nil {
		fmt.Printf("Fehler beim Dekodieren der Datei: %v\n", err)
		return
	}
	index--

	// Entferne den Eintrag im Slice
	slice = slices.Delete(slice, index, index+1)

	// Setze den Dateizeiger zurück und lösche den aktuellen Inhalt der Datei
	err = file.Truncate(0)
	if err != nil {
		fmt.Printf("Fehler beim Zurücksetzen des Dateizeigers: %v\n", err)
		return
	}

	// Setze den Dateizeiger auf den Anfang der Datei zurück
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Printf("Fehler beim Zurücksetzen des Dateizeigers auf den Anfang der Datei: %v\n", err)
		return
	}

	// Schreibe das aktualisierte Slice in die Datei
	err = json.NewEncoder(file).Encode(slice)
	if err != nil {
		fmt.Printf("Fehler beim Schreiben des aktualisierten Slices in die Datei: %v\n", err)
		return
	}

	fmt.Println("Eintrag wurde erfolgreich gelöscht.")
}
