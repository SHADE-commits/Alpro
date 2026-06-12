package main

const maxVoter = 100

type Voter struct {
	NIM      string
	HasVoted bool
}

var voters [maxVoter]Voter
var voterCount int
