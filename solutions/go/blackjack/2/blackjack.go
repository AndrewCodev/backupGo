package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var valueCard int

    switch {
    case card == "ace":
        valueCard = 11
    case card == "two":
        valueCard = 2
    case card == "three":
        valueCard = 3
	case card == "four":
        valueCard = 4
	case card == "five":
        valueCard = 5
	case card == "six":
        valueCard = 6
	case card == "seven":
        valueCard = 7
	case card == "eight":
        valueCard = 8
	case card == "nine":
        valueCard = 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king" :
        valueCard = 10
    default:
        valueCard = 0
    }
    return valueCard
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var valueCard1 int = ParseCard(card1)
    var valueCard2 int = ParseCard(card2)
    var sumValCards int = valueCard1 + valueCard2
    var valueDealerCard int = ParseCard(dealerCard)
	var firstTurn string  

    switch {
        case sumValCards == 22 :
        	firstTurn = "P"
        case sumValCards == 21 && valueDealerCard < 10 :
			firstTurn = "W"
        case (sumValCards >= 17 && sumValCards <= 20) ||
        	(sumValCards >= 12 && sumValCards <= 16 && valueDealerCard < 7) ||
        	(sumValCards == 21 && valueDealerCard >= 10): 
			firstTurn = "S"
        case (sumValCards >= 12 && sumValCards <= 16 && valueDealerCard >= 7) || sumValCards <= 11:
			firstTurn = "H"
    }
    return firstTurn
}
