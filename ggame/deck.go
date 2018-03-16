package ggame

import (
	"fmt"
	"strings"
)

type Deck struct {
	Cards     []Card
	NoOfDecks int
}

func NewDeck() *Deck {
	d := Deck{Cards: make([]Card, 52), NoOfDecks: 1}
	index := 0
	for s := range SUITS {
		for r := range RANKS {
			d.Cards[index] = Card(r*4 + s)
			fmt.Println(index, ":", r*4+s)
			index++
		}
	}
	return &d
}

func (d *Deck) Print() {
	index := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			fmt.Println(d.Cards[index].String())
			index++
		}
	}
}

func (d *Deck) ToString() string {
	var buffer strings.Builder
	index := 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			buffer.WriteString(d.Cards[index].String())
			buffer.WriteString(" ")
			index++
		}
	}
	return buffer.String()
}

func (d *Deck) Shuffle() {

}

func (d *Deck) CardsRemaiing() int {
	return len(d.Cards)
}

// Not sure which name would be good..so keep two functions
func (d *Deck) NoOfCards() int {
	return len(d.Cards)
}

func (d *Deck) Take() Card {
	var c Card
	if len(d.Cards) > 0 {
		c = d.Cards[0]
		d.Cards = d.Cards[1:]
	}
	return c
}
