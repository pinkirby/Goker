package Goker

type Suit int8

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

var SuitStringRep map[Suit]string = map[Suit]string{
	Clubs:    "clubs",
	Diamonds: "diamonds",
	Hearts:   "hearts",
	Spades:   "spades",
}

type Card struct {
	powerVal    int8
	straightVal int8
	fullName    string
	smallRep    string
	suit        Suit
}

func compareCards(firstCard Card, secondCard Card) int8 {
	if firstCard.powerVal < firstCard.powerVal {
		return -1
	} else if firstCard.powerVal > firstCard.powerVal {
		return 1
	}
	return 0
}
