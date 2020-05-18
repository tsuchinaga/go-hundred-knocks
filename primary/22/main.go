package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if n <= -10 || 10 <= n {
		fmt.Println("OK")
	}
}
