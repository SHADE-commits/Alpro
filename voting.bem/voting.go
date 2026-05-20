package main

import "fmt"

func vote() {
	var nim string

	fmt.Print("Enter Your NIM: ")
	fmt.Scan(&nim)

	index := findVoter(nim)

	if index == -1 {

		voters = append(voters, Voter{
			NIM:      nim,
			HasVoted: false,
		})

		index = len(voters) - 1
	}

	if voters[index].HasVoted {
		fmt.Println("You have already voted")
		return
	}

	showCandidates()

	fmt.Print("Choose candidate number: ")

	var choice int
	fmt.Scan(&choice)

	if choice <= 0 || choice > len(candidates) {
		fmt.Println("Invalid candidate")
		return
	}

	candidates[choice-1].Votes++
	voters[index].HasVoted = true

	fmt.Println("Vote submitted successfully")
}
