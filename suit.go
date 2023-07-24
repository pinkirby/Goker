package goker

type Suit struct {
	shortRep string
	fullName string
}

var Clubs Suit = Suit{"C", "clubs"}
var Diamonds Suit = Suit{"D", "diamonds"}
var Hearts Suit = Suit{"H", "hearts"}
var Spades Suit = Suit{"S", "spades"}

var Suits map[string]*Suit = map[string]*Suit{
	"C": &Clubs,
	"D": &Diamonds,
	"H": &Hearts,
	"S": &Spades,
}
