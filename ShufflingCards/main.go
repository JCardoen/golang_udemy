package main

func main() {
	deckOfCards := deckFromFile("cards.csv")
	deckOfCards.shuffle()
	deckOfCards.print()
}
