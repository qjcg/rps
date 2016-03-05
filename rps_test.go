package rps

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestRandomHand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	v := RandomHand()
	vt := reflect.TypeOf(v)
	if vt.Kind() != reflect.Int || vt.Name() != "Hand" {
		t.Error("expected a Hand, got", v)
	}
}

func TestPlay(t *testing.T) {
	var testcases = []struct {
		p1, p2   Hand
		expected Result
	}{
		// tie
		{Rock, Rock, Tie},
		{Paper, Paper, Tie},
		{Scissors, Scissors, Tie},

		// p1 wins
		{Rock, Scissors, WinP1},
		{Paper, Rock, WinP1},
		{Scissors, Paper, WinP1},

		// p2 wins
		{Rock, Paper, WinP2},
		{Paper, Scissors, WinP2},
		{Scissors, Rock, WinP2},
	}

	for _, c := range testcases {
		r := Play(c.p1, c.p2)
		if r != c.expected {
			t.Errorf("expected: %v, got: %v", c.expected, r)
		}
	}
}
