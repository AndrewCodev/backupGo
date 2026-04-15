package blackjack

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
	case "ten", "jack", "queen", "king":
        return 10
    default:
        return 0
    }
}

func FirstTurn(card1, card2, dealerCard string) string {
    sumValCards := ParseCard(card1) + ParseCard(card2)
    valueDealerCard := ParseCard(dealerCard)

    switch {
    case sumValCards == 22:
        return "P"

    case sumValCards == 21 && valueDealerCard < 10:
        return "W"

    case sumValCards == 21:
        return "S"

    case sumValCards >= 17 && sumValCards <= 20:
        return "S"

    case sumValCards >= 12 && valueDealerCard >= 7:
        return "H"

    case sumValCards >= 12:
        return "S"

    default:
        return "H"
    }
}
