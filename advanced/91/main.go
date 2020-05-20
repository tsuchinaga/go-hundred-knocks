package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// デックを作る
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

	// シャッフル
	rand.Seed(time.Now().UnixNano())
	for i := range cards {
		desc := int(rand.Float64() * 52)
		cards[i], cards[desc] = cards[desc], cards[i]
	}

	// 1枚手元にある状態から1枚ずつ手札に加えていく
	dc := []card{cards[0]}
	for i := 1; i < len(cards); i++ {
		dc = append(dc, cards[i])
		total := sum(dc)
		output(dc, total)
		if total < 17 {
			continue
		} else if total < 22 {
			fmt.Println("これでOKです")
			break
		} else {
			fmt.Println("バストです")
			break
		}
	}
}

// card - カード
type card struct {
	suit   int    // 0: spade, 1: heart, 2: club, 3: diamond
	number int    // 1-13
	str    string // A, 2-10, J, Q, K
}

// sum - カードの合計をブラックジャックのルールに則って計算する
func sum(cards []card) int {
	// 1-13までの数字カードの枚数
	nums := make([]int, 13)
	for _, c := range cards {
		nums[c.number-1]++
	}

	sum := 0
	for i := 12; i >= 0; i-- {
		switch i {
		case 12, 11, 10:
			sum += 10 * nums[i]
		case 0:
			for j := nums[0]; j > 0; j-- {
				if sum+11*j <= 21 {
					sum += 11 * j
					break
				}
				sum += 1
			}
		default:
			sum += (i + 1) * nums[i]
		}
	}
	return sum
}

// output - 課題の形で標準出力
func output(cards []card, total int) {
	s := make([]string, len(cards))
	for i, c := range cards {
		s[i] = c.str
	}
	fmt.Printf("%s : 合計%d\n", strings.Join(s, " "), total)
}
