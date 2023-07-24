package goker

import (
	"fmt"
	"strconv"
)

type Category struct {
	shortRep string
	fullName string
}

var Ace Category = Category{"A", "Ace"}
var King Category = Category{"K", "King"}
var Queen Category = Category{"Q", "Queen"}
var Jack Category = Category{"J", "Jack"}
var Ten Category = Category{"T", "10"}

var SpecialCategories map[string]*Category = make(map[string]*Category, 13)

func initCategories() {
	SpecialCategories["A"] = &Ace
	SpecialCategories["K"] = &King
	SpecialCategories["Q"] = &Queen
	SpecialCategories["J"] = &Jack
	SpecialCategories["T"] = &Ten
	for i := 2; i < 10; i++ {
		num := strconv.Itoa(i)
		SpecialCategories[num] = &Category{num, num}
	}
	for _, elem := range SpecialCategories {
		fmt.Println(elem)
	}
}
