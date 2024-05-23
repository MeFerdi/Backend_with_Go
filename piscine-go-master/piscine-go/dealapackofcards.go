package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	numCards := 12
	numPlayers := 4

	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Player %d: ", i+1)
		for j := 0; j < numCards/numPlayers; j++ {
			index := i*numCards/numPlayers + j
			if j > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			fmt.Printf("%d", deck[index])
		}
		z01.PrintRune('\n')
	}
}
