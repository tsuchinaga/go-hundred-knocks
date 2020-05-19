package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	i := 0
	for {
		if n == 1 {
			break
		}

		i++
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		fmt.Printf("%d: %d\n", i, n)
	}
}
