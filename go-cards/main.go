package main

func main() {
	game := Game{}
	game.start()

	game.showDeck()
	game.showHands()

	game.giveCardNextPlayer()
	game.showHands()

	go game.send_message("hola")
	game.print_message()

}
