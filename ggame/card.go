package ggame

import (
	"fmt"
)

type Suit int // Spade, Heart, Diamond, Club
type Rank int // Two,Three,Four,Five,Six,Seven,Eight,Nine,Ten,Jack,Queen,King,Ace - Increasing order

const (
	CLUB Suit = iota
	DIAMOND
	HEART
	SPADE
)

const (
	TWO Rank = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

type Card int

var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	RANKS = []Rank{TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING, ACE}
)

func (c Card) String() string {
	r := Rank(c / 4)
	s := Suit(c % 4)

	rank := ""
	suit := ""

	switch r {
	case ACE:
		rank = "A"
	case TWO:
		rank = "2"
	case THREE:
		rank = "3"

	case FOUR:
		rank = "4"

	case FIVE:
		rank = "5"

	case SIX:
		rank = "6"

	case SEVEN:
		rank = "7"

	case EIGHT:
		rank = "8"

	case NINE:
		rank = "9"

	case TEN:
		rank = "10"

	case JACK:
		rank = "J"

	case QUEEN:
		rank = "Q"

	case KING:
		rank = "K"

	}

	switch s {
	case CLUB:
		suit = "C"
	case DIAMOND:
		suit = "D"
	case HEART:
		suit = "H"
	case SPADE:
		suit = "S"
	}
	return fmt.Sprintf("%s%s", rank, suit)

}
