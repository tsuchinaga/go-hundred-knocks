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
	printCard(card)
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
