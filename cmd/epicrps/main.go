package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/qjcg/rps"

	"github.com/mgutz/ansi"
)

const (
	RockAscii     = "0"
	PaperAscii    = "#"
	ScissorsAscii = "X"

	RockLetter     = "R"
	PaperLetter    = "P"
	ScissorsLetter = "S"

	Scissors     = "✂"
	ScissorsHand = "✌"
)

type ScoreBoard struct {
	p1      int
	p2      int
	ties    int
	games   int
	pctP1   float64
	pctP2   float64
	pctTies float64
}

func (s *ScoreBoard) Percentages() {
	s.pctP1 = 100.0 * float64(s.p1) / float64(s.games)
	s.pctP2 = 100.0 * float64(s.p2) / float64(s.games)
	s.pctTies = 100.0 * float64(s.ties) / float64(s.games)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	HandSymbols := map[int]string{
		rps.Rock:     RockAscii,
		rps.Paper:    PaperAscii,
		rps.Scissors: Scissors,
	}

	ngames := flag.Int("n", 100000, "number of games to play")
	live := flag.Bool("l", false, "display live scoreboard updates")
	flag.Parse()

	paleBlue := ansi.ColorFunc("69+b:black")

	// Set up a new scoreboard and play!
	sb := ScoreBoard{}
	for {
		// Exit when ngames have been played
		if sb.games == *ngames {
			// only print scoreboard if no live updates (otherwise redundant)
			if *live {
				fmt.Println()
			} else {
				sb.Percentages()
				fmt.Printf("\rp1: %v (%.3f%%)  p2: %v (%.3f%%)  ties: %v (%.3f%%)  games: %v/%d\n", sb.p1, sb.pctP1, sb.p2, sb.pctP2, sb.ties, sb.pctTies, sb.games, *ngames)
			}
			os.Exit(0)
		}

		handP1 := rps.RandomHand()
		handP2 := rps.RandomHand()

		switch rps.Play(handP1, handP2) {
		case rps.Tie:
			sb.ties++
		case rps.WinP1:
			sb.p1++
		case rps.WinP2:
			sb.p2++
		}

		sb.games++
		if *live {
			sb.Percentages()
			fmt.Printf("\r(%.1f%%) %v  %s %s  %v (%.1f%%)  ties: %v (%.1f%%)  games: %v/%d ", sb.pctP1, sb.p1, paleBlue(HandSymbols[handP1]), paleBlue(HandSymbols[handP2]), sb.p2, sb.pctP2, sb.ties, sb.pctTies, sb.games, *ngames)
		}
	}
}
