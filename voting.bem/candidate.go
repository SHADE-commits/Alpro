package main

import "fmt"

type Candidate struct {
	Name  string
	Votes int
}

var candidates []Candidate

func addCandidate() {
	var name string

	fmt.Print("Candidate name: ")
	fmt.Scan(&name)

	candidates = append(candidates, Candidate{
		Name:  name,
		Votes: 0,
	})

	fmt.Println("Candidate added successfully")
}

func showCandidates() {
	fmt.Println("\n=== Candidate List ===")

	for i, c := range candidates {
		fmt.Println(i+1, "-", c.Name, "| Votes:", c.Votes)
	}
}
