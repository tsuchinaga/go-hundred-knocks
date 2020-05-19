package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("石の数を入力してください（10以上）: ")
	_, _ = fmt.Scan(&n)
	fmt.Printf("石の数: %d\n", n)

	p := 0 // 0: プレイヤー, 1: コンピュータ
	if (n-1)%4 == 0 {
		fmt.Println("あなたからどうぞ")
	} else {
		fmt.Println("ではわたしから")
		p = 1
	}

	for {
		if p == 1 {
			fmt.Printf("%d個取ります\n", (n-1)%4)
			n -= (n - 1) % 4
		} else {
			fmt.Print("何個取る（1〜3個）? ")
			_, _ = fmt.Scan(&m)
			if m < 1 || 3 < m {
				continue
			}
			n -= m
		}
		fmt.Printf("石の数: %d\n", n)

		if n <= 1 {
			break
		}
		p = (p + 1) % 2
	}

	fmt.Println("わたしの勝ち")
}
