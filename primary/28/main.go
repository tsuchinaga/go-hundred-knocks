package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	factorial := 1
	for i := 2; i <= n; i++ {
		factorial *= i
	}
	fmt.Printf("factorial = %d\n", factorial)
}
