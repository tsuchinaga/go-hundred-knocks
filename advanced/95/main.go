package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := ""
	for i := 0; i < 4; i++ {
		target += strconv.Itoa(int(rand.Float64() * 10))
	}

	var s string
	fmt.Print("4桁の数字を入力: ")
	_, _ = fmt.Scan(&s)

	fmt.Printf("target = %s\n", target)

	var hit, blow int
	used := make([]int, 4)

	// hitを確認
	for i := range target {
		if target[i] == s[i] {
			used[i] = 1
			hit++
		}
	}

	// blowを確認
	for i, t := range target {
		if used[i] != 0 {
			continue
		}
		for _, c := range s {
			if t == c {
				used[i] = 1
				blow++
				break
			}
		}
	}

	fmt.Printf("%d hit, %d blow\n", hit, blow)
}
