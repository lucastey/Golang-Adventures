package main

func main() {
	//var card string = "Ace of Spades"
	//card := newCard()
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")

	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
