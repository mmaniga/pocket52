package ggame

import "testing"

var valueTests = []struct {
	in  string
	out HandRank
}{
	{"2D 4S 5D 6D 7D", HighCard},
	{"4C 4S 5D 6D 7D", OnePair},
	{"4C 4S 5D 5H 7D", TwoPair},
	{"4C 4S 4D 5H 7D", ThreeOfAKind},
	{"2D 3S 4D 5D 6D", Straight},
	{"2D 4D 5D 6D 7D", Flush},
	{"2D 3D 4D 5D 6D", StraightFlush},
	{"TD JD QD KD AD", RoyalFlush},
}

func TestHand_GetHandRank(t *testing.T) {
	for _, h := range valueTests {
		tmphand, _ := NewHandFromString(h.in)
		handRank := tmphand.GetHandRank()
		if handRank != h.out {
			t.Errorf("hand:%s Failure did not produce expected value(got %s expected %s)", h.in, handRank.ToString(), h.out.ToString())
		} else {
			t.Errorf("hand:%s  Success Got expected value(got %s expected %s)", h.in, handRank.ToString(), h.out.ToString())
		}
	}
}
