package main

import (
	"fmt"
	"github.com/pocket52/ggame"
)

func main() {
	fmt.Println("Testing Pocket52")
	d := ggame.NewDeck()
	fmt.Println(d.ToString())
	d.Shuffle()
	fmt.Println(d.ToString())

	h := ggame.Hand{}
	cards, _ := d.PopN(5)
	h.FiveCard(cards)
	fmt.Println(h.ToString())

}
