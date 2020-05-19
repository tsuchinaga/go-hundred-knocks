package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	// fmt.Printf("%016b\n", n) // これでできるけど、一応配列でやってみる

	ary := make([]int, 16)
	for i := 15; i >= 0; i-- {
		ary[i] = n % 2
		n /= 2
	}
	for _, b := range ary {
		fmt.Print(b)
	}
	fmt.Println()
}
