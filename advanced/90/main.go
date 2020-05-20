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
		var s string
		switch i%13 + 1 {
		case 1:
			s = "A"
		case 11:
			s = "J"
		case 12:
			s = "Q"
		case 13:
			s = "K"
		default:
			s = strconv.Itoa(i%13 + 1)
		}
		cards[i] = card{suit: i / 13, number: i%13 + 1, str: s}
	}
	for i := range cards {
		desc := int(rand.Float64() * 52)
		cards[i], cards[desc] = cards[desc], cards[i]
	}

	sum := 0
	for i := 0; i < 2; i++ {
		switch cards[i].number {
		case 1:
			sum += 11
		case 11, 12, 13:
			sum += 10
		default:
			sum += cards[i].number
		}
	}

	fmt.Printf("%s %s : 合計%d\n", cards[0].str, cards[1].str, sum)
}

type card struct {
	suit   int    // 0: spade, 1: heart, 2: club, 3: diamond
	number int    // 1-13
	str    string // A, 2-10, J, Q, K
}
