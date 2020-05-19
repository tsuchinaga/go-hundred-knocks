package main

import "fmt"

func main() {
	var n int
	fmt.Print("距離 ")
	_, _ = fmt.Scan(&n)

	sum := 610
	n -= 1700
	for n > 0 {
		sum += 80
		n -= 313
	}
	fmt.Printf("金額: %d\n", sum)
}
