package main

type Player struct {
	Hand Deck
	Name string
}

func (p Player) showHand() {
	p.Hand.showDeck()
}

func (p *Player) giveCard(next_p *Player) {
	p.Hand.deal(1, next_p)
}
