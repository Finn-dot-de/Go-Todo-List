package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"todo/assets/structs"
)

func Done(index int) {
	// Öffne die JSON-Datei
	file, err := os.OpenFile("todo_list.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Dekodiere den Inhalt der Datei in ein Slice
	var slice []structs.TodoList
	err = json.NewDecoder(file).Decode(&slice)
	if err != nil {
		panic(err)
	}
	index--

	// Entferne den Eintrag im Slice
	slice = slices.Delete(slice, index, index+1)

	// Setze den Dateizeiger zurück und lösche den aktuellen Inhalt der Datei
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// Setze den Dateizeiger auf den Anfang der Datei zurück
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	// Schreibe das aktualisierte Slice in die Datei
	err = json.NewEncoder(file).Encode(slice)
	if err != nil {
		panic(err)
	}

	fmt.Println("Eintrag wurde erfolgreich gelöscht.")
}
