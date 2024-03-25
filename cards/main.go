package main

import "fmt"

// Array: fixed length list
// Slice: variable length list

func main() {
	// var card string = "Ace of Spades" --> Same as below
	cards := newDeck()
	cards= append(cards, "Six of Spades")

	hand, remainingCards := deal(cards, 3)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("cards_test")
}
