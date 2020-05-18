package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	s := make([]string, 0)
	for i := 0; i < n/5; i++ {
		s = append(s, "*****")
	}
	if n%5 > 0 {
		s = append(s, strings.Repeat("*", n%5))
	}
	fmt.Println(strings.Join(s, " "))
}
