package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if 5 < n && n < 20 {
		fmt.Println("OK")
	}
}
