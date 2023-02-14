package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/fs"
	"os"
)

const HAND_SIZE = 7
const DECK_SIZE = 52
const PLAYERS_NUMBER = 4

type Game struct {
	Session_id string
	*Deck
	Players [PLAYERS_NUMBER]*Player
}

func (g *Game) start() {
	g.Session_id = "Session 1"
	g.Deck = &Deck{}
	i := 0
	for i < PLAYERS_NUMBER {
		g.Players[i] = &Player{}
		i += 1
	}

	fmt.Println("Creating deck")
	g.Deck.newDeck()
	fmt.Printf("\nDealing initial hand to %v players", PLAYERS_NUMBER)
	g.deal()
}

func (g *Game) deal() {
	for _, player := range g.Players {
		go g.Deck.deal(HAND_SIZE, player)
	}
}

// go binary encoder
func (g Game) encode() []byte {
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	err := e.Encode(g)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}
	return b.Bytes()
}

// go binary decoder
func (g *Game) decode(str string) {
	by, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println(`failed base64 Decode`, err)
	}
	b := &bytes.Buffer{}
	b.Write(by)
	d := gob.NewDecoder(b)
	err = d.Decode(g)
	if err != nil {
		fmt.Println(`failed gob Decode`, err)
	}
}

func (g Game) saveToFile(filename string, filepath string) string {
	saved_state := g.encode()
	var anyone_can_read_write fs.FileMode = 0666
	os.WriteFile(filepath+filename, saved_state, anyone_can_read_write)
	return base64.StdEncoding.EncodeToString(saved_state)
}

func (g *Game) loadFromFile(filepath string) {
	saved_state, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Println(`failed load from file`, err)
	}

	(*g).loadFromToken(base64.StdEncoding.EncodeToString(saved_state))
}

func (g *Game) loadFromToken(token string) {
	(*g).decode(token)
}
