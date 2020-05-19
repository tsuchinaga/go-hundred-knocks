package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("input money: ")
	_, _ = fmt.Scan(&n)

	ary := make([]int, 3)
	ary[0] = n / 100
	n %= 100
	ary[1] = n / 10
	ary[2] = n % 10

	fmt.Printf("100円玉%d枚, 10円玉%d枚, 1円玉%d枚\n", ary[0], ary[1], ary[2])
}
