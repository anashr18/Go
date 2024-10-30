package main

func main() {
	// cards := deck{"Ace of Spades", "King of Hearts", "Queen of Diamonds"}
	// cards = append(cards, "Jack of Clubs", newCard())

	// cards := newDeck()
	// cards.print()
	// hand, rem := deal(cards, 5)
	// hand.print()
	// rem.print()
	// greetings := "HEllo world!!!"
	// fmt.Println([]byte(greetings))
	// fmt.Println(cards.toString())
	// cards.saveToFile("cards")
	deck_disk, _ := readFile("cards")
	deck_disk.print()
	deck_disk.shuffle()
	deck_disk.print()
}

func newCard() string {
	return "Ace of Diamonds"
}
