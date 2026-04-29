package main

import (
	"fmt"
)

func main() {

	type Entry struct {
		chore string
		done  bool
	}

	var list []Entry

	/* 
		var arrayList [5]Entry
		fixed length array, not dynamic. We want to be able to add as many chores as we want, so we use a slice instead of an array.
	*/

	addToList := func(e Entry) {
		list = append(list, e)
	}

	showList := func() {
		for i, entry := range list {
			fmt.Printf("%d: %+v\n ", i, entry)
		}
	}


	toggleEntry := func(e *Entry) {
		e.done = !e.done
	}

	var userOption int

	printMenu := func() {
		fmt.Println("1. Add a chore")
		fmt.Println("2. Toggle a chore")
		fmt.Println("3. Show list")
		fmt.Println("4. Exit")
	}

	play := func() {	
		fmt.Scanln(&userOption)

		switch userOption {
		case 1:
			var chore string
			fmt.Print("Enter a chore: ")
			fmt.Scanln(&chore)
			addToList(Entry{chore: chore, done: false})
		case 2:
			var index int
			fmt.Print("Enter the index of the chore to toggle: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(list) {
				toggleEntry(&list[index])
			} else {
				fmt.Println("Invalid index")
			}
		case 3:
			showList()
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}

	for userOption != 4 {
		printMenu()
		play()
	}
}