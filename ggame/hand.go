package ggame

import (
	"errors"
	//"fmt"
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

func (h HandRank) ToString() string {
	return HandRenkString[h]
}

var (
	HandsRank      = []HandRank{HighCard, OnePair, TwoPair, ThreeOfAKind, Straight, Flush, FullHouse, FourOfAKind, StraightFlush, RoyalFlush}
	HandRenkString = []string{"HighCard", "OnePair", "TwoPair", "ThreeOfAKind", "Straight", "Flush", "FullHouse", "FourOfAKind", "StraightFlush", "RoyalFlush"}
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

func NewHandFromString(hand string) (*Hand, error) {

	h := Hand{}
	hs := strings.Split(hand, " ")
	if len(hs) != 5 {
		return &h, errors.New("invalid hand")
	}
	cards := make([]Card, 5)
	for i := 0; i < 5; i++ {
		card, _ := NewCard(hs[i])
		cards[i] = card
	}
	h.FiveCard(cards)
	return &h, nil
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

	var handRank HandRank
	straight := true

	for i := 0; i < len(h.Cards)-1; i++ {
		if h.Cards[i].Rank()+1 != h.Cards[i+1].Rank() {
			straight = false
			break
		}
	}

	flush := true
	for i := 1; i < len(h.Cards); i++ {
		if h.Cards[0].Suit() != h.Cards[i].Suit() {
			flush = false
			break
		}
	}

	if straight {
		handRank = Straight
	}

	if flush {
		handRank = Flush
	}

	if straight && flush {
		handRank = StraightFlush
	}

	return handRank
}
