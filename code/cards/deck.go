package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cards := deck{}

	var template string

	for _, suit := range suits {
		template = " of " + suit

		for _, value := range values {
			cards = append(cards, value+template)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}
