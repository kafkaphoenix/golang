package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value string
}

type Deck []Card

func (d Deck) showDeck() {
	fmt.Println()
	for i, card := range d {
		fmt.Printf("%v %v of %v\n", i+1, card.Value, card.Suit)
	}
}

func (d *Deck) newDeck() {
	suits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	values := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
	for _, suit := range suits {
		for _, value := range values {
			*d = append(*d, Card{Suit: suit, Value: value})
		}
	}
	d.shuffle()
}

func (d *Deck) deal(amount int, p *Player) {
	cards := (*d)[:amount]
	*d = (*d)[amount:]
	p.Hand = append(p.Hand, cards...)
}

func (d *Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range *d {
		newPosition := r.Intn(len(*d) - 1)
		(*d)[i], (*d)[newPosition] = (*d)[newPosition], (*d)[i]
	}
}
