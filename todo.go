package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"todo/assets/utils"
)

const (
	errorArgument = `Wrong argument !!
	
Use "todo help" for more information about a command.`

	helpComment = `todo add  <description>  add new task
todo del  <number>       delete task
todo done <number>       finish task
todo help                print help
todo list                list help`
)

var uebergabe string = ""

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errorArgument)
	}
	var auswahl_command string = os.Args[1]
	if len(os.Args) > 2 {

		uebergabe = strings.Join(os.Args[2:], " ")

		fmt.Printf("Übergabe: %v", uebergabe)
		menue(auswahl_command, uebergabe)

	} else {
		menue(auswahl_command, uebergabe)
	}

}

func menue(auswahl string, uebergabe string) {
	switch auswahl {
	case "add":
		utils.AddEntry(uebergabe)
	case "del":
		i, err := strconv.Atoi(uebergabe)
		if err != nil {
			panic(err)
		}
		utils.Delete(i)
	case "done":
		i, err := strconv.Atoi(uebergabe)
		if err != nil {
			panic(err)
		}
		utils.Done(i)
	default:
		switch auswahl {
		case "list":
			tasks, err := utils.ReadJsonFile()
			if err != nil {
				log.Fatal(err)
			}
			// Ausgabe der gelesenen Aufgabenliste
			utils.PrintTasks(tasks)
		case "help":
			fmt.Println(helpComment)
		default:
			log.Fatal(errorArgument)

		}

	}
}
