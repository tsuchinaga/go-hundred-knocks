package main

import (
	"fmt"
)

func main() {
	min, max := 1, 99

	var n, a int
	for {
		a++
		fmt.Printf("%dですか? ", min+(max-min)/2)
		_, _ = fmt.Scan(&n)

		if n == 0 {
			break
		} else if n < 0 {
			max = min + (max-min)/2 - 1
		} else {
			min = min + (max-min)/2 + 1
		}
	}
	fmt.Printf("%d回であてました\n", a)
}
