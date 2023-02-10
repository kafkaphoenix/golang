package main

func main() {
	game := Game{}
	game.start()

	game.Deck.ShowDeck()
	game.Player.ShowHand()
}
