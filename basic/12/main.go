package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("Hello World!")
	}
}
