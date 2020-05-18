package main

import "fmt"

func main() {
	fmt.Print("input number: ")

	var n int
	_, _ = fmt.Scan(&n)

	fmt.Printf("answer = %d\n", n*3)
}
