package ggame

import "fmt"

type Deck struct {
	Cards []Card
	NoOfDecks int
}

func NewDeck() (*Deck) {
	d:= Deck{Cards:make([]Card,52),NoOfDecks:1}
	index := 0
	for s:= range SUITS {
		for r := range RANKS {
			d.Cards[index] = Card(s*4+r)
			index++
		}
	}
	return &d
 }


func (d *Deck) Print()  {
	index :=0
	for i := 0; i< 4;i++ {
		for j :=0;j<13;j++ {
			fmt.Println(d.Cards[index])
			index++
		}
	}
}