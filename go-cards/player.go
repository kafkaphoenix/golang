package main

type Player struct {
	Hand Deck
}

func (p Player) ShowHand() {
	p.Hand.ShowDeck()
}
