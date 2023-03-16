package cards

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func GenerateDeck() Deck {
	deck := Deck{}
	for suit := Spades; suit <= Clubs; suit++ {
		for rank := Ace; rank <= King; rank++ {
			card := GenerateCard(suit, rank)
			deck.Cards = append(deck.Cards, card)
		}
	}
	return deck
}

func (d *Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	seed := rand.New(source)
	for i := range d.Cards {
		j := seed.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]

	}
}

func (d *Deck) Deal(n int) []Card {
	cards := d.Cards[:n]
	d.Cards = d.Cards[:n]
	return cards
}

func (d Deck) Len() int {
	return len(d.Cards)
}

func (d Deck) IsEmpty() bool {
	return len(d.Cards) == 0
}
