package main

func main() {
	cards := newDeckFromFile("my_cards.txt")

	if (cards == deck{}) {
		cards = randomDeck()
		cards.saveToFile("my_cards_txt")
	}

	cards.shuffle()
	cards.print()
}
