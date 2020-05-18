package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if n > 0 {
		fmt.Println("positive")
	} else if n < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
}
