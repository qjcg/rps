// An library for playing the classic game of Rock, Paper, Scissors, intended as a learning tool.
package rps

import (
	"math/rand"
)

type (
	Hand   int
	Result int
)

const (
	// Hands
	Rock Hand = iota
	Paper
	Scissors

	// Results
	P1Win Result = iota + 10
	P2Win
	Tie
)

// RandomHand returns a random Hand.
func RandomHand() Hand {
	return [3]Hand{Rock, Paper, Scissors}[rand.Intn(3)]
}

// Play a single game of Rock, Paper, Scissors and return the result.
func Play(p1, p2 Hand) Result {
	var r Result
	switch {
	case p1 == p2:
		r = Tie
	case (p1 == Rock && p2 != Paper),
		(p1 == Paper && p2 != Scissors),
		(p1 == Scissors && p2 != Rock):
		r = P1Win
	default:
		r = P2Win
	}
	return r
}
