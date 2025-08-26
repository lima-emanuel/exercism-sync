package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	v1 := ParseCard(card1)
	v2 := ParseCard(card2)
	vd := ParseCard(dealerCard)
	hand_sum := v1 + v2

	switch {
	case card1 == "ace" && card1 == card2:
		return "P"
	case hand_sum == 21 && vd < 10:
		return "W"
	case hand_sum == 21 && vd >= 10:
		return "S"
	case hand_sum >= 17 && hand_sum <= 20:
		return "S"
	case hand_sum >= 12 && hand_sum <= 16 && vd < 7:
		return "S"
	case hand_sum >= 12 && hand_sum <= 16 && vd >= 7:
		return "H"
	case hand_sum <= 11:
		return "H"
	default:
		return ""
	}
}
