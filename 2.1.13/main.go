package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Deck sort. Explain how you would put a deck of cards in order by suit (in the order
// spades, hearts, clubs, diamonds) and by rank within each suit, with the restriction
// that the cards must be laid out face down in a row, and the only allowed operations
// are to check the values of two cards and to exchange two cards (keeping them face down).

// One way is to just turn two random cards that we know aren't in the correct position.
// If one card if fx. 3 of hearts, we know that the position is 1*13 + 3 -> the index of the suit
// times 13, plus the rank of the card. So after we turn two cards, we can place them exactly where they need to be.

// In my implementation i use a shell sort algorithm starten with a gap of 18 halving every time.
// Every loop i look at two cards and compare their absolute position in the deck,
// and swap them if one position is bigger than the other.

const (
	spades = iota
	hearts
	clubs
	diamonds
)

type card struct {
	suit int
	rank int
}

type deck []*card

func main() {
	d := getDeck()

	rand.Seed(time.Now().Unix())
	// https://github.com/golang/go/wiki/SliceTricks#shuffling
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}

	fmt.Println("Shuffled Deck:")
	d.print()

	d.sort()

	fmt.Println()
	fmt.Println("Sorted Deck:")
	d.print()
}

func (d deck) sort() {
	gap := 18

	for {
		sorted := d.sortRank(gap)
		if sorted {
			break
		}

		if gap > 1 {
			gap /= 2
		}
	}
}

func (d deck) sortRank(gap int) bool {
	for i := 0; i+gap < len(d); i++ {
		if d[i].suit*13+d[i].rank > d[i+gap].suit*13+d[i+gap].rank {
			// fmt.Println("swapping", i, i+gap, d[i], d[i+gap], d[i].suit*13+d[i].rank, d[i+gap].suit*13+d[i+gap].rank)
			d[i], d[i+gap] = d[i+gap], d[i]
		}
	}

	rankSorted := true
	for i := 0; i < len(d)-1; i++ {
		if d[i].suit*13+d[i].rank > d[i+1].suit*13+d[i+1].rank {
			rankSorted = false
		}
	}
	return rankSorted
}

func getDeck() deck {
	d := make(deck, 0)	
	for suit := 0; suit < 4; suit++ {	
		for i := 1; i <= 13; i++ {
			d = append(d, &card{suit, i})
		}
	}
	return d
}

// func (d deck) Len() int {
// 	return len(d)
// }

// func (d deck) Swap(i, j int) {
// 	d[i], d[j] = d[j], d[i]
// }

// func (d deck) Less(i, j int) bool {
// 	return d[i].suit*13+d[i].rank < d[j].suit*13+d[j].rank
// }

//printing

func (d deck) print() {
	for i := range d {
		fmt.Println(d[i])
	}
}

func (c *card) String() string {
	suit := getSuitString(c.suit)
	rank := getRankString(c.rank)
	return fmt.Sprintf("%s of %s", rank, suit)
}

func getSuitString(suit int) string {
	switch suit {
	case spades:
		return "spades"
	case hearts:
		return "hearts"
	case clubs:
		return "clubs"
	case diamonds:
		return "diamonds"
	}
	panic("unknown suit")
}

func getRankString(rank int) string {
	switch rank {
	case 1:
		return "Ace"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return strconv.Itoa(rank)
	}
}
