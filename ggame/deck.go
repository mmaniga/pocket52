package ggame

import (
	"errors"
	"fmt"
	"math/rand"
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

	totalCards := len(d.Cards)

	for i := 0; i < totalCards; i++ {
		r := i + rand.Intn(totalCards-i)
		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
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

// Not sure of which name to keep for the function..
func (d *Deck) Pop() Card {
	return d.Take()
}

func (d *Deck) PopN(nCards int) ([]Card, error) {
	cards := make([]Card, nCards)

	if len(d.Cards) < nCards {
		return cards, errors.New("Less cards in Deck")
	}
	for i := 0; i < nCards; i++ {
		cards[i] = d.Pop()
	}
	return cards, nil

}
