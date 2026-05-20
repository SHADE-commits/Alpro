package main

import "fmt"

func sortResults() {
	for i := 1; i < len(candidates); i++ {
		key := candidates[i]
		j := i - 1

		for j >= 0 && candidates[j].Votes < key.Votes {
			candidates[j+1] = candidates[j]
			j--
		}

		candidates[j+1] = key
	}
}

func showResults() {
	sortResults()

	fmt.Println("\n===== Voting Results =====")

	totalVotes := 0

	for _, c := range candidates {
		totalVotes += c.Votes
	}

	for _, c := range candidates {
		percentage := 0.0

		if totalVotes > 0 {
			percentage = float64(c.Votes) / float64(totalVotes) * 100
		}

		fmt.Printf("%s - %d votes (%.2f%%)\n",
			c.Name,
			c.Votes,
			percentage)
	}

	if len(candidates) > 0 {
		fmt.Println("Winner:", candidates[0].Name)
	}
}
