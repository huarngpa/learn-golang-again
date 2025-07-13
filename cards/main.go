package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard() // Equivalent to var card string = "Ace of Spades"

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades") // This actually creates a new slice and returns it

	fmt.Printf("card=%v\n", card)

	for i, card := range cards {
		fmt.Printf("i=%v. card=%v\n", i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
