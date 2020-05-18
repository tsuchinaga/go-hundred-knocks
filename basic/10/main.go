package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if n < 0 {
		n *= -1
	}
	fmt.Printf("absolute value is %d\n", n)
}
