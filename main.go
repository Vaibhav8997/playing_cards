package main

// import "fmt"
func main() {
	// //var card string = "Ace of spades"
	// card := newCard() //colun equal to operator
	// fmt.Println(card)

	//slice
	// cards := newDeck()

	// //cards.print()
	// hand, remaningCards := deal(cards, 5)

	// hand.print()
	// remaningCards.print()

	cards := newDeck()
	cards.saveToFile("my_caards")

}
