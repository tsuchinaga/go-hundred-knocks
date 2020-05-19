package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ans := int(rand.Float64()*99) + 1

	var n, a int
	for {
		a++
		fmt.Print("数を入力: ")
		_, _ = fmt.Scan(&n)

		if n == ans {
			break
		} else if n < ans {
			fmt.Println("それより大きいです")
		} else {
			fmt.Println("それより小さいです")
		}
	}
	fmt.Printf("正解!! %d回でクリア\n", a)
}
