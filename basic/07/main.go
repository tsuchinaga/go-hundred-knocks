package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if n == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("not zero")
	}
}
