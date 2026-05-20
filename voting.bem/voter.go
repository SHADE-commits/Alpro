package main

import "fmt"

type Voter struct {
	NIM      string
	Name     string
	HasVoted bool
}

var voters []Voter

func addVoter() {
	var nim, name string

	fmt.Print("Enter NIM: ")
	fmt.Scan(&nim)

	fmt.Print("Enter Candidate Name: ")
	fmt.Scan(&name)

	voters = append(voters, Voter{
		NIM:      nim,
		Name:     name,
		HasVoted: false,
	})

	fmt.Println("Voter added successfully")
}

func findVoter(nim string) int {
	for i, v := range voters {
		if v.NIM == nim {
			return i
		}
	}

	return -1
}
