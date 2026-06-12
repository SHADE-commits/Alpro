package main

import "fmt"

func vote() {
	var nim string

	fmt.Print("Enter Your NIM: ")
	fmt.Scan(&nim)

	index := sequentialSearchVoter(nim)

	if index == -1 {
		voters[voterCount].NIM = nim
		voters[voterCount].HasVoted = false

		index = voterCount
		voterCount++
	}

	if voters[index].HasVoted {
		fmt.Println("You have already voted!")
		return
	}

	showCandidates()

	var choice int

	fmt.Print("Choose candidate number: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > candidateCount {
		fmt.Println("Invalid candidate")
		return
	}

	candidates[choice-1].Votes++
	voters[index].HasVoted = true

	fmt.Println("Vote submitted successfully")
}
