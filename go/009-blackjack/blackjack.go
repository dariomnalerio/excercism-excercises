package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var action string
	sum := ParseCard(card1) + ParseCard(card2)

	switch {
	case sum == 22:
		action = "P"
	case sum == 21:
		if ParseCard(dealerCard) == 11 || ParseCard(dealerCard) == 10 {
			action = "S"
		} else {
			action = "W"
		}
	case sum >= 17 && sum <= 20:
		action = "S"
	case sum >= 12 && sum <= 16:
		if ParseCard(dealerCard) >= 7 {
			action = "H"
		} else {
			action = "S"
		}
	case sum <= 11:
		action = "H"
	}

	return action
}
