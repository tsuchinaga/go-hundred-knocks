package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("石の数を入力してください（10以上）: ")
	_, _ = fmt.Scan(&n)

	p := 0
	for {
		fmt.Printf("プレイヤー%dの番です\n", p+1)
		fmt.Print("何個取る（1〜3個）? ")
		_, _ = fmt.Scan(&m)
		if m < 1 || 3 < m {
			continue
		}

		n -= m
		fmt.Printf("石の数: %d\n", n)

		if n == 1 {
			fmt.Printf("プレイヤー%dの勝ち\n", p+1)
			break
		} else if n <= 0 {
			fmt.Printf("プレイヤー%dの反則負け\n", p+1)
			break
		}
		p = (p + 1) % 2
	}
}
