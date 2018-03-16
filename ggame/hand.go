package ggame

import (
	"strings"
)

/**
High Card: Highest value card.
One Pair: Two cards of the same value.
Two Pairs: Two different pairs.
Three of a Kind: Three cards of the same value.
Straight: All cards are consecutive values.
Flush: All cards of the same suit.
Full House: Three of a kind and a pair.
Four of a Kind: Four cards of the same value.
Straight Flush: All cards are consecutive values of same suit.
Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
*/

type HandRank int

const (
	HIGH_CARD HandRank = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_OF_KIND
	STRAIGHT_FLUSH
	ROYAL_FLUSH
)

var (
	HANDSRank = []HandRank{HIGH_CARD, ONE_PAIR, TWO_PAIR, THREE_OF_A_KIND, STRAIGHT, FLUSH, FULL_HOUSE, FOUR_OF_KIND, STRAIGHT_FLUSH, ROYAL_FLUSH}
)

type Hand struct {
	Cards []Card
	Desc  string
	Rank  HandRank
}

func (h *Hand) FiveCard(c []Card) {
	h.Cards = c
}

func (h *Hand) ToString() string {
	var buffer strings.Builder
	noOfCards := len(h.Cards)
	for i := 0; i < noOfCards; i++ {
		buffer.WriteString(h.Cards[i].String())
		buffer.WriteString(" ")
	}
	return buffer.String()
}
