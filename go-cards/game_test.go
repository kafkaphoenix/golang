package main

import (
	"os"
	"testing"
)

func TestSaveAndLoadGame(t *testing.T) {
	os.Remove("_decktesting")

	game := Game{}
	game.start()
	game.saveToFile("_decktesting", "./")
	game.Deck = &Deck{}

	if len(*game.Deck) != 0 {
		t.Errorf("Expected loaded game has a wrong deck length %v", len(*game.Deck))
	}

	game.loadFromFile("./_decktesting")

	if len(*game.Deck) != 24 {
		t.Errorf("Expected loaded game has a wrong deck length %v", len(*game.Deck))
	}

	os.Remove("_decktesting")
}

func TestSaveAndLoadGameToken(t *testing.T) {
	os.Remove("_decktesting")

	game := Game{}
	game.start()
	token := game.saveToFile("_decktesting", "./")
	game.Deck = &Deck{}

	if len(*game.Deck) != 0 {
		t.Errorf("Expected loaded game has a wrong deck length %v", len(*game.Deck))
	}

	game.loadFromToken(token)

	if len(*game.Deck) != 24 {
		t.Errorf("Expected loaded game has a wrong deck length %v", len(*game.Deck))
	}

	os.Remove("_decktesting")
}
