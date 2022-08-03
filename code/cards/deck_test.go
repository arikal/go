package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	actualSize := len(d)
	expectedSize := 52

	if actualSize != expectedSize {
		t.Errorf("Expected deck length of %v, but got %v", expectedSize, actualSize)
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected Ace of Clubs as first card, but got %v", d[0])
	}

	if d[actualSize-1] != "King of Spades" {
		t.Errorf("Expected King of Spades as first card, but got %v", d[actualSize-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove(".deck_testing")

	deck := newDeck()
	deck.saveToFile(".deck_testing")

	loadedDeck := newDeckFromFile(".deck_testing")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected %v cards in loaded deck, got %v", len(deck), len(loadedDeck))
	}
}
