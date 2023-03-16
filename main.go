package main

import (
	"fmt"
	"poker-engine/pkg/cards"
)

func main() {
	deck := cards.GenerateDeck()
	deck.Shuffle()
	fmt.Println(deck, "yo")
}
