package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if n < -10 {
		fmt.Println("range 1")
	} else if n < 0 {
		fmt.Println("range 2")
	} else {
		fmt.Println("range 3")
	}
}
