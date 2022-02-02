package main

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of spades"
	// card = newCard()
	// card = "Five of Diamond"
	// cards := newDeck()
	// cards = append(cards, "Hello, world")

	// cards1, _ := deal(cards, 4)

	// cards1.print()
	// cards2.print()

	// cards.saveToFile("test.txt")

	// fmt.Println(cards.toString())

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards[2])
	cards := newDeckFromFile("test.txt")
	// fmt.Println(newDeckFromFile("test.txt"))
	cards.shuffle()
	cards.print()
}

// func newCard() string {
// 	return "Ace of spades"
// }
