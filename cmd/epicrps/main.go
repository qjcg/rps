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
	ROCK_ASCII     = "0"
	PAPER_ASCII    = "#"
	SCISSORS_ASCII = "X"

	ROCK_LETTER     = "R"
	PAPER_LETTER    = "P"
	SCISSORS_LETTER = "S"

	SCISSORS      = "✂"
	SCISSORS_HAND = "✌"
)

type ScoreBoard struct {
	p1      int
	p2      int
	ties    int
	pctP1   float64
	pctP2   float64
	pctTies float64
	games   int
}

func (s *ScoreBoard) Percentages() {
	s.pctP1 = 100.0 * float64(s.p1) / float64(s.games)
	s.pctP2 = 100.0 * float64(s.p2) / float64(s.games)
	s.pctTies = 100.0 * float64(s.ties) / float64(s.games)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	HandSymbols := map[rps.Hand]string{
		rps.Rock:     ROCK_ASCII,
		rps.Paper:    PAPER_ASCII,
		rps.Scissors: SCISSORS,
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

		p1_hand := rps.RandomHand()
		p2_hand := rps.RandomHand()

		switch rps.Play(p1_hand, p2_hand) {
		case rps.Tie:
			sb.ties++
		case rps.P1Win:
			sb.p1++
		case rps.P2Win:
			sb.p2++
		}

		sb.games++
		if *live {
			sb.Percentages()
			fmt.Printf("\r(%.1f%%) %v  %s %s  %v (%.1f%%)  ties: %v (%.1f%%)  games: %v/%d ", sb.pctP1, sb.p1, paleBlue(HandSymbols[p1_hand]), paleBlue(HandSymbols[p2_hand]), sb.p2, sb.pctP2, sb.ties, sb.pctTies, sb.games, *ngames)
		}
	}

	var input string
	fmt.Scanln(&input)
}
