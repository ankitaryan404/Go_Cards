package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Length
	l := len(d)
	if l != 16 {
		t.Errorf("Expected deck of length %d, but got %v", 16, l)
	}

	// First card
	fc := "Ace of Spades"
	if d[0] != fc {
		t.Errorf("Expected first card to be %v, but got %v", fc, d[0])
	}

	// Last card
	lc := "Four of Clubs"
	if d[len(d)-1] != lc {
		t.Errorf("Expected last card to be %v, but got %v", lc, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	// Length
	l := len(loadedDeck)
	if l != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", l)
	}

	os.Remove("_decktesting")
}