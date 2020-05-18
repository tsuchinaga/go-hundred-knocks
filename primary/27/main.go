package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)
}
