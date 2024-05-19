package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cards = map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}

	v, ok := cards[card]
	if !ok {
		return 0
	}

	return v
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1v := ParseCard(card1)
	c2v := ParseCard(card2)
	dc := ParseCard(dealerCard)

	switch {
	case c1v == 11 && c2v == 11:
		return "P"
	case c1v+c2v == 21:
		if dc == 11 || dc == 10 {
			return "S"
		}
		return "W"
	case c1v+c2v >= 17 && c1v+c2v <= 20:
		return "S"
	case c1v+c2v >= 12 && c1v+c2v <= 16:
		if dc >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
