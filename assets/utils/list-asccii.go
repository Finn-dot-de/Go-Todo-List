package utils // Paket "utils"

import ( // Importiert die benötigten Pakete
	"encoding/json"       // Paket zum Codieren und Decodieren von JSON
	"fmt"                 // Paket für das Formatieren von Zeichenfolgen
	"os"                  // Paket zum Arbeiten mit dem Betriebssystem
	"strings"             // Paket für Operationen mit Zeichenfolgen
	"todo/assets/structs" // Importiert die Struktur "TodoList" aus einem lokalen Paket
)

// max gibt das Maximum von zwei ganzen Zahlen zurück
func max(a, b int) int {
	if a > b { // Wenn a größer als b ist
		return a // Gib a zurück
	}
	return b // Sonst gib b zurück
}

// calculateColumnWidths berechnet die maximale Breite jeder Spalte in der ASCII-Tabelle
func calculateColumnWidths(tasks []structs.TodoList) (int, int, int) {
	maxCreated := len("Erstellt")  // Länge des Texts "Erstellt"
	maxDone := len("Erledigt")     // Länge des Texts "Erledigt"
	maxDesc := len("Beschreibung") // Länge des Texts "Beschreibung"

	for _, task := range tasks { // Für jede Aufgabe in den Aufgaben
		maxCreated = max(maxCreated, len(task.Created)) // Aktualisiere maxCreated mit der Länge des Erstellungsdatums der aktuellen Aufgabe
		maxDone = max(maxDone, len(task.Done))          // Aktualisiere maxDone mit der Länge des Erledigungsdatums der aktuellen Aufgabe
		maxDesc = max(maxDesc, len(task.Desc))          // Aktualisiere maxDesc mit der Länge der Beschreibung der aktuellen Aufgabe
	}

	return maxCreated, maxDone, maxDesc // Gib die maximalen Breiten zurück
}

// PrintTasksAsccii druckt die Aufgaben als ASCII-Tabelle
func PrintTasksAscii(tasks []structs.TodoList) {
	maxCreated, maxDone, maxDesc := calculateColumnWidths(tasks) // Berechne die maximalen Spaltenbreiten

	// Header
	header := fmt.Sprintf("+-%s-+-%s-+-%s-+", strings.Repeat("-", maxCreated), strings.Repeat("-", maxDone), strings.Repeat("-", maxDesc))      // Erzeuge den Tabellenkopf
	fmt.Println(header)                                                                                                                         // Drucke den Tabellenkopf
	fmt.Printf("| %-3s | %-6s | %-*s | %-*s | %-*s |\n", "Nr.", "Status", maxCreated, "Erstellt", maxDone, "Erledigt", maxDesc, "Beschreibung") // Drucke die Spaltenüberschriften
	fmt.Println(header)                                                                                                                         // Drucke den Tabellenkopf erneut

	// Rows
	for i, task := range tasks { // Für jede Aufgabe und ihren Index
		Check := "[ ]"       // Standardmäßig ist die Checkbox unausgefüllt
		if task.Done != "" { // Wenn die Aufgabe als erledigt markiert ist
			Check = "[X]" // Ändere die Checkbox auf ausgefüllt
		}
		fmt.Printf("| %-3d | %-6s | %-*s | %-*s | %-*s |\n", i+1, Check, maxCreated, task.Created, maxDone, task.Done, maxDesc, task.Desc) // Drucke die Zeile der Tabelle für die aktuelle Aufgabe
	}

	// Footer
	fmt.Println(header) // Drucke den Tabellenfuß
}

// ReadJsonFileAsccii liest Aufgaben aus einer JSON-Datei und gibt sie als Slice von structs.TodoList zurück
func ReadJsonFileAscii() ([]structs.TodoList, error) {
	file, err := os.OpenFile("todo_list.json", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0644) // Öffne die JSON-Datei zum Lesen
	if err != nil {                                                                       // Wenn ein Fehler auftritt
		return nil, fmt.Errorf("fehler beim öffnen der datei: %s", err) // Gib eine Fehlermeldung zurück
	}
	defer file.Close() // Stelle sicher, dass die Datei am Ende geschlossen wird

	dekoder := json.NewDecoder(file) // Erzeuge einen JSON-Decoder für die Datei

	var aufgaben []structs.TodoList                   // Erzeuge ein Slice von structs.TodoList für die Aufgaben
	if err := dekoder.Decode(&aufgaben); err != nil { // Versuche, die JSON-Daten in die Aufgaben zu decodieren
		return nil, fmt.Errorf("fehler beim lesen der JSON-daten: %s", err) // Gib eine Fehlermeldung zurück, wenn ein Fehler auftritt
	}

	return aufgaben, nil // Gib die Aufgaben zurück
}
