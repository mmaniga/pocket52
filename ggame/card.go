package ggame

type Suit int  // Spade, Heart, Diamond, Club
type Rank int  // Two,Three,Four,Five,Six,Seven,Eight,Nine,Ten,Jack,Queen,King,Ace - Increasing order


const (
	CLUB Suit  = iota
	DIAMOND
	HEART
	SPADE
)

const (
	ACE Rank = iota+1
	TWO
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

)

type Card int

var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	RANKS = []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)

