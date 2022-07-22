package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//filename := "my_hand"
	//cards.print()
	//hand, _ := deal(cards, 5)

	//err := hand.saveToFile(filename)
	//if err != nil {
	//	return
	//}

	//fmt.Println("Deck Saved")

	//fmt.Println("trying to open de deck file")

	//newDeck := newDeckFromFile(filename)

	//fmt.Println(newDeck.toString())

}
