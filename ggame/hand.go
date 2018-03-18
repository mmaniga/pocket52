package ggame

import (
	"sort"
	"strings"
)

/**
1) High Card: Highest value card.
2) One Pair: Two cards of the same value.
3) Two Pairs: Two different pairs.
4) Three of a Kind: Three cards of the same value.
5) Straight: All cards are consecutive values.
6) Flush: All cards of the same suit.
7) Full House: Three of a kind and a pair.
8) Four of a Kind: Four cards of the same value.
9) Straight Flush: All cards are consecutive values of same suit.
10) Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
*/

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var (
	HandsRank = []HandRank{HighCard, OnePair, TwoPair, ThreeOfAKind, Straight, Flush, FullHouse, FourOfAKind, StraightFlush, RoyalFlush}
)

type Hand struct {
	Cards []Card
	Desc  string
	Rank  HandRank
}

func (h *Hand) Len() int           { return len(h.Cards) }
func (h *Hand) Swap(i, j int)      { h.Cards[i], h.Cards[j] = h.Cards[j], h.Cards[i] }
func (h *Hand) Less(i, j int) bool { return h.Cards[i].Rank() < h.Cards[j].Rank() }

func (h *Hand) Sort() {
	sort.Sort(h)
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

func (h *Hand) GetHandRank() HandRank {

	return HighCard
}
