package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected lenght of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDEckAndNewDeckFromFile(t *testing.T) {

	deleteTestingFiles := func() {
		if _, err := os.Stat("_decktesting"); err == nil {
			err := os.Remove("_decktesting")
			if err != nil {
				t.Errorf("The testing files could not be deleted: %v", err)
			}
		}
	}

	deleteTestingFiles()

	deck := newDeck()
	err := deck.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("The file could not be created:%v", err)
	}

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	deleteTestingFiles()
}
