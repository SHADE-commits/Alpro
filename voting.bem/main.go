package main

import "fmt"

func main() {
	for {
		fmt.Println("\n===== E-Vote BEM =====")
		fmt.Println("1. Add Candidate")
		fmt.Println("2. Vote")
		fmt.Println("3. Show Results")
		fmt.Println("4. Exit")
		fmt.Print("Choose menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addCandidate()

		case 2:
			vote()

		case 3:
			showResults()

		case 4:
			fmt.Println("Program ended")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
