package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	cards := make([]card, 52)
	for i := range cards {
		cards[i] = card{suit: i / 13, number: i%13 + 1}
	}
	for i := range cards {
		desc := int(rand.Float64() * 52)
		cards[i], cards[desc] = cards[desc], cards[i]
	}

	for _, c := range cards {
		s := ""
		switch c.suit {
		case 0:
			s = "スペード"
		case 1:
			s = "ハート"
		case 2:
			s = "クローバー"
		case 3:
			s = "ダイヤ"
		}
		n := ""
		switch c.number {
		case 1:
			n = "A"
		case 2, 3, 4, 5, 6, 7, 8, 9, 10:
			n = strconv.Itoa(c.number)
		case 11:
			n = "J"
		case 12:
			n = "Q"
		case 13:
			n = "K"
		}
		fmt.Printf("%s%s\n", s, n)
	}
}

type card struct {
	suit   int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // 1-13
}
