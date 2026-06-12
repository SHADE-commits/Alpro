package main

func sequentialSearchVoter(nim string) int {
	for i := 0; i < voterCount; i++ {
		if voters[i].NIM == nim {
			return i
		}
	}

	return -1
}

func binarySearchCandidate(name string) int {
	left := 0
	right := candidateCount - 1

	for left <= right {
		mid := (left + right) / 2

		if candidates[mid].Name == name {
			return mid
		}

		if name > candidates[mid].Name {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
