package main

func main() {
	// var card string = "Ace of Spades"
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	//  fmt.Println(cards)
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))

	// cards := newDeck()

	//string(file) hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }
