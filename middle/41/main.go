package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	not := " not"
	if 0 < n && n < 10 {
		not = ""
	}
	fmt.Printf("%d is%s a single figure.\n", n, not)
}
