package goker

import (
	"testing"
)

type CardTemplate struct {
	*Card
	expected bool
}

func TestCardValiditiy(t *testing.T) {
	//Arrange
	var cards = []CardTemplate{
		{&Card{1, 1, "2 of Spades", "2S", &Spades}, true},
		{&Card{13, 0, "Ace of Hearts", "AH", &Hearts}, true},
		{&Card{13, 0, "Ace of Hearts", "AH", &Spades}, false},
		{&Card{0, 0, "Ace of Hearts", "AH", &Hearts}, false},
		{&Card{13, 13, "Ace of Hearts", "AH", &Spades}, false},
		{&Card{10, 10, "Jack of Clubs", "JC", &Clubs}, true},
		{&Card{10, 10, "Jack of Clubs", "JH", &Clubs}, false},
		{&Card{10, 10, "Jack of Clubs", "YC", &Clubs}, false},
		{&Card{10, 10, "Jack of Clubs", "JQ", &Clubs}, false},
		{&Card{10, 10, "Jack of Clubs", "JP", &Clubs}, false},
		{&Card{10, 10, "Jack of Clubs", "JP", &Diamonds}, false},
		{&Card{10, 10, "Jack Clubs", "JP", &Diamonds}, false},
		{&Card{10, 10, "Jack arstarstastar Clubs", "JC", &Clubs}, false},
		{&Card{8, 10, "Jack of Clubs", "JC", &Clubs}, true},
		{&Card{-1, 10, "Jack of Clubs", "JC", &Clubs}, false},
		{&Card{14, 10, "Jack of Clubs", "JC", &Clubs}, false},
		{&Card{10, 14, "Jack of Clubs", "JC", &Clubs}, false},
		{&Card{11, 11, "Queen of Diamonds", "QD", &Diamonds}, true},
	}

	//Act
	//Assert
	for i, card := range cards {
		if card.IsValid() != card.expected {
			t.Fatalf(
				"the card at index %d: %s validity should be %t",
				i,
				card.ToString(),
				card.expected)
		}
	}

}
