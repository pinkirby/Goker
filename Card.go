package goker

import "fmt"

type Card struct {
	powerVal    int8
	straightVal int8
	fullName    string
	shortRep    string
	suit        *Suit
	category    string
}

func (*Card) IsValid() bool {
	return true
}

func (c *Card) ToString() string {
	return fmt.Sprintf("%s", c.fullName)
}

func compareCards(firstCard Card, secondCard Card) int8 {
	if firstCard.powerVal < firstCard.powerVal {
		return -1
	} else if firstCard.powerVal > firstCard.powerVal {
		return 1
	}
	return 0
}
