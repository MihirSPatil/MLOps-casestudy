package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var value int
	switch card{
        default:
        	value = 0
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
    }
    return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer. S,H,P,W
func FirstTurn(card1, card2, dealerCard string) string {
	var action string
    
    if card1 == card2 && card1 == "ace"{
        action = "P"
    }
    
    cardSum := ParseCard(card1) + ParseCard(card2)
    dealerValue := ParseCard(dealerCard)
    
    switch {
        case cardSum == 21:
        	if dealerValue < 10{
        		action = "W"
            } else{
                action = "S"
            }
        case 17 <= cardSum && cardSum <= 20:
        	action = "S"
        case 12 <= cardSum && cardSum <=16:
        	if dealerValue >= 7{
                action = "H"
            } else {action = "S"
                   }
        case cardSum <= 11:
        	action = "H"
    }
    return action
}
