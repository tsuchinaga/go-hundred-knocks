package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}
