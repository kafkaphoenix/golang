package main

func main() {
	game := Game{}
	game.start()

	game.Deck.ShowDeck()
	for _, player := range game.Players {
		player.ShowHand()
	}
}
