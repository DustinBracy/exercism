package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	cards := map[string]int {
		"ace": 11,
		"king": 10,
		"queen": 10,
		"jack": 10,
		"ten": 10,
		"nine": 9,
		"eight": 8,
		"seven":7,
		"six": 6,
		"five": 5,
		"four": 4,
		"three": 3,
		"two": 2,
	}
	return cards[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var action string
	c1 := ParseCard(card1)
	c2 := ParseCard(card2)
	dealer := ParseCard(dealerCard)
	switch {
		case card1 == "ace" && card2 == "ace":
			action = "P"
		case c1 + c2 == 21:
			if dealer < 10 {
				action = "W"
			} else {
				action = "S"
			}
		case c1 + c2 > 16:
			action = "S"
		case c1 + c2 > 11:
			if dealer > 6 {
				action = "H"
			} else {
				action = "S"
			}
		case c1 + c2 < 12:
			action = "H"
		}
	return action
}
