package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("input 1st value: ")
	_, _ = fmt.Scan(&n)
	fmt.Print("input 2nd value: ")
	_, _ = fmt.Scan(&m)

	l := n / m
	fmt.Printf("result: %d\n", l)
	fmt.Printf("result: %d\n", l*m)
}
