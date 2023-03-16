package cards

import "fmt"

type Suit int

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Rank int

const (
	Ace Rank = iota + 1 // index from 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// represents standard playing card with a suit and a rank
type Card struct {
	Suit Suit
	Rank Rank
}

func GenerateCard(suit Suit, rank Rank) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

func (c Card) String() string {
	var rankStr string
	switch c.Rank {
	case Ace:
		rankStr = "Ace"
	case Jack:
		rankStr = "Jack"
	case Queen:
		rankStr = "Queen"
	case King:
		rankStr = "King"
	default:
		rankStr = fmt.Sprintf("%d", c.Rank)
	}
	var suitStr string
	switch c.Suit {
	case Spades:
		suitStr = "Spades"
	case Hearts:
		suitStr = "Hearts"
	case Diamonds:
		suitStr = "Diamonds"
	case Clubs:
		suitStr = "Clubs"
	}
	return fmt.Sprintf("%s of %s", rankStr, suitStr)
}
