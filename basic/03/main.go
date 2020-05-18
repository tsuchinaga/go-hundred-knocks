package main

import "fmt"

func main() {
	fmt.Print("input number: ")

	var n int
	_, _ = fmt.Scan(&n)

	fmt.Printf("your number is %d\n", n)
}
