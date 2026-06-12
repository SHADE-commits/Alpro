package main

import "fmt"

func main() {
	var menu int

	for menu != 5 {
		fmt.Println("\n===== E-VOTE BEM =====")
		fmt.Println("1. Add Candidate")
		fmt.Println("2. Delete Candidate")
		fmt.Println("3. Vote")
		fmt.Println("4. Show Result")
		fmt.Println("5. Exit")

		fmt.Print("Choose menu: ")
		fmt.Scan(&menu)

		switch menu {
		case 1:
			addCandidate()

		case 2:
			deleteCandidate()

		case 3:
			vote()

		case 4:
			insertionSortDescending()
			showCandidates()

		case 5:
			fmt.Println("Program finished")

		default:
			fmt.Println("Invalid menu")
		}
	}
}
