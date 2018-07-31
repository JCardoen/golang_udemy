package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()
	if len(cards) != 52 {
		t.Errorf("Expected length 52 but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", cards[0])
	}

	if cards[51] != "Jack of Clubs" {
		t.Errorf("Expected first card to be Jack of Clubs but got %v", cards[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("deck_testing.csv")

	d := newDeck()
	d.toFile("deck_testing")

	loadedDeck := deckFromFile("deck_testing.csv")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(loadedDeck))
	}

	if d[0] != loadedDeck[0] {
		t.Errorf("Expected first cards of both decks to be the same but got: %v, %v", d[0], loadedDeck[0])
	}
}