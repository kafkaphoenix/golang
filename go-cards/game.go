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

type Game struct {
	Session_id string
	*Deck
	*Player
}

func (g *Game) start() {
	g.Session_id = "Session 1"
	g.Deck = &Deck{}
	g.Player = &Player{}

	fmt.Println("Creating deck")
	g.Deck.newDeck()
	fmt.Println("Dealing initial hand")
	g.Deck.deal(HAND_SIZE, g.Player)
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
