package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	win := []int{0, 0}
	for i := 0; i < 5; i++ {
		for {
			fmt.Print("あなたの手を選んでください(グー0、チョキ1、パー2): ")
			var n int
			_, _ = fmt.Scan(&n)
			p := int(rand.Float64() * 3)
			switch p {
			case 0:
				fmt.Println("コンピュータはグー")
			case 1:
				fmt.Println("コンピュータはチョキ")
			case 2:
				fmt.Println("コンピュータはパー")
			}

			switch n - p {
			case -1, 2:
				fmt.Println("あなたの勝ち")
				win[0]++
			case 0:
				fmt.Println("あいこ")
				continue
			case -2, 1:
				fmt.Println("わたしの勝ち")
				win[1]++
			default:
				fmt.Println("そんな手はありません。あなたの負け")
				win[1]++
			}
			break
		}
		fmt.Printf("あなた%d勝、わたし%d勝\n", win[0], win[1])
	}

	if win[0] >= 3 {
		fmt.Println("あなたの総合勝利です。参りました。")
	}
}
