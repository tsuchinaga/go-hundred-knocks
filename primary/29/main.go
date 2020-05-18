package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 5; i++ {
		var n int
		fmt.Print("input number: ")
		_, _ = fmt.Scan(&n)
		sum += n
	}
	fmt.Printf("sum = %d\n", sum)
}
