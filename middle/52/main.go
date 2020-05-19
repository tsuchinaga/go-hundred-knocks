package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("input year: ")
	_, _ = fmt.Scan(&n)

	yn := "ない"
	if n%4 == 0 && (n%100 != 0 || n%400 == 0) {
		yn = "ある"
	}

	fmt.Printf("%d年は閏年で%s\n", n, yn)
}
