package main

import "fmt"

const maxCandidate = 20

type Candidate struct {
	Name  string
	Votes int
}

var candidates [maxCandidate]Candidate
var candidateCount int

func addCandidate() {
	var name string

	fmt.Print("Enter candidate name: ")
	fmt.Scan(&name)

	candidates[candidateCount].Name = name
	candidates[candidateCount].Votes = 0
	candidateCount++

	fmt.Println("Candidate added successfully")
}

func showCandidates() {
	fmt.Println("\n=== Candidate List ===")

	for i := 0; i < candidateCount; i++ {
		fmt.Println(i+1, "-", candidates[i].Name, "| Votes:", candidates[i].Votes)
	}
}

func deleteCandidate() {
	var number int

	showCandidates()

	fmt.Print("Enter candidate number to delete: ")
	fmt.Scan(&number)

	for i := number - 1; i < candidateCount-1; i++ {
		candidates[i] = candidates[i+1]
	}

	candidateCount--

	fmt.Println("Candidate deleted successfully")
}
