package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)
	fmt.Println(strings.Repeat("*", n))
}
