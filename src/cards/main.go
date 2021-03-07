package main

import "fmt"

func main() {

	cards := []string{newCard(), "Ace of diamonds"}
	cards = append(cards, "Six of spades")

	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of diamonds"
}
