package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	for i := 0; i <= n; i += 2 {
		fmt.Println(i)
	}
}
