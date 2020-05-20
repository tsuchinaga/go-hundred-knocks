package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	numbers := make([]int, 75)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// シャッフル
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100_000; i++ {
		j := int(rand.Float64() * 75)
		numbers[i%75], numbers[j] = numbers[j], numbers[i%75]
	}

	var card [5][5]int
	for i := 0; i < 25; i++ {
		if i == 12 {
			continue // 真ん中は0のままにする
		}
		card[i/5][i%5] = numbers[i]
	}

	var n, ans int
	for {
		ans++
		printCard(card)
		fmt.Printf("%d個目: ", ans)
		_, _ = fmt.Scan(&n)

		i, j := search(card, n)
		if 0 <= i && 0 <= j {
			card[i][j] = 0
			if isBingo(card) {
				printCard(card)
				fmt.Println("***** BINGO *****")
				break
			} else {
				fmt.Println("あった！")
			}
		}
	}
}

func printCard(card [5][5]int) {
	for _, l := range card {
		s := make([]string, 5)
		for j, n := range l {
			s[j] = strconv.Itoa(n)
		}
		fmt.Println(strings.Join(s, "\t"))
	}
}

func search(card [5][5]int, n int) (int, int) {
	for i, l := range card {
		for j, m := range l {
			if n == m {
				return i, j
			}
		}
	}
	return -1, -1
}

func isBingo(card [5][5]int) bool {
	// 横
	for _, l := range card {
		b := true
		for _, n := range l {
			if n != 0 {
				b = false
				break
			}
		}
		if b {
			return true
		}
	}

	// 縦
	for i := 0; i < 5; i++ {
		b := true
		for j := 0; j < 5; j++ {
			if card[j][i] != 0 {
				b = false
				break
			}
		}
		if b {
			return true
		}
	}

	// 斜め
	if (card[0][0]+card[1][1]+card[2][2]+card[3][3]+card[4][4] == 0) || (card[0][4]+card[1][3]+card[2][2]+card[3][1]+card[4][0] == 0) {
		return true
	}

	return false
}
