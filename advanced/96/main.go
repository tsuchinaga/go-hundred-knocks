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
	// fmt.Printf("target = %s\n", target)
	// target = "3169"

	ans := 0
	for {
		ans++
		fmt.Printf("%d回目\n", ans)

		var s string
		fmt.Print("4桁の数字を入力: ")
		_, _ = fmt.Scan(&s)

		var hit, blow int
		used := make([]int, 4)

		// hitを確認
		for i := range target {
			if target[i] == s[i] {
				used[i] = 1
				hit++
			}
		}

		if hit == 4 {
			break
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
	fmt.Println("正解")
}
