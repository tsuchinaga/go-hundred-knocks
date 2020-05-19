package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	oe := "odd"
	if n%2 == 0 {
		oe = "even"
	}
	fmt.Printf("%d is %s.\n", n, oe)
}
