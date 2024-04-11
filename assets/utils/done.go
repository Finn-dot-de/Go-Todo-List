package utils

import (
	"encoding/json"
	"fmt"
	"todo/assets/filethings"
)

func Done(index int) {
	tasks, file := filethings.OpenAndReadFile()
	if file != nil {
		defer file.Close()
	} else {
		return
	}

	index--

	// Ändere den Wert "Done" im Slice
	tasks[index].Done = todayDate()

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

	// Schreibe das aktualisierte Slice in die Datei
	err := json.NewEncoder(file).Encode(tasks)
	if err != nil {
		fmt.Println("Fehler beim Schreiben in die Datei:", err)
		return
	}

	fmt.Println("Eintrag wurde erfolgreich aktualisiert.")
}
