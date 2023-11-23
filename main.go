package main

func main() {
	// One - Create a new deck and save it as my_cards
	cards := newDeck()
	// cards.saveToFile("my_cards") // Save a deck in a file

	// Two - Create a new deck from the file my_cards
	// cards := newDeckFromFile("my_cards")
	// cards := newDeckFromFile("error_case") // Error case
	cards.shuffle()
	cards.print()
}