package main

import "fmt"

func main() {
	for {
		var n int
		fmt.Print("input number: ")
		_, _ = fmt.Scan(&n)
		if n == 0 {
			break
		}
	}
}
