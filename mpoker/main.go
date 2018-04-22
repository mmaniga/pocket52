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
	d.Shuffle()
	fmt.Println(d.ToString())
	d.Shuffle()
	fmt.Println(d.ToString())

	h := ggame.Hand{}
	cards, _ := d.PopN(5)
	h.FiveCard(cards)
	fmt.Println(h.ToString())
	h.Sort()
	fmt.Println(h.ToString())

	h2 := ggame.Hand{}
	c2, _ := d.PopN(5)
	h2.FiveCard(c2)
	fmt.Println(h2.ToString())
	h2.Sort()
	fmt.Println(h2.ToString())

	h3, _ := ggame.NewHandFromString("2C 3C 4C 5C 6C")
	fmt.Println(h3.ToString())
	fmt.Println(h3.GetHandRank().ToString())

	h4, _ := ggame.NewHandFromString("2C 3C 4D 5C 6C")
	fmt.Println(h4.ToString())
	fmt.Println(h4.GetHandRank().ToString())

	h5, _ := ggame.NewHandFromString("2C 3C 7D 5C 6C")
	fmt.Println(h5.ToString())
	fmt.Println(h5.GetHandRank().ToString())

	h6, _ := ggame.NewHandFromString("2C 7C TC AC JC")
	fmt.Println(h6.ToString())
	fmt.Println(h6.GetHandRank().ToString())

}
