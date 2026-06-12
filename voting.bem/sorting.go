package main

func insertionSortDescending() {
	for i := 1; i < candidateCount; i++ {

		key := candidates[i]
		j := i - 1

		for j >= 0 && candidates[j].Votes < key.Votes {
			candidates[j+1] = candidates[j]
			j--
		}

		candidates[j+1] = key
	}
}
