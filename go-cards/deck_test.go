package main

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := Deck{}
	deck.newDeck()
	if len(deck) != DECK_SIZE {
		t.Errorf("Expected deck length of 52, but got %v", len(deck))
	}
}

func TestDeal(t *testing.T) {
	deck := Deck{}
	player := &Player{}
	c := make(chan string)
	deck.newDeck()
	go deck.deal(7, player, c)
	fmt.Println(<-c)

	if len(deck) != 45 {
		t.Errorf("Expected deck length of 45, but got %v", len(deck))
	}

	if len(player.Hand) != 7 {
		t.Errorf("Expected player hand length of 7, but got %v", len(player.Hand))
	}
}
