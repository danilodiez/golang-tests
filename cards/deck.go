package main

import "fmt"

// Create a new type of a 'deck'
// Which is a slice of strings

type deck []string

// Receiver function
// any var of type deck now have access to that function
// as d.print()
// for convention we call receiver vars with one letter
func (d deck) print()  {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {	
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+ " of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {

	// from init to handSize, from handSize to end
	return d[:handSize], d[handSize:]

}

