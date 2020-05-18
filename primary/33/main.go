package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	for i := 1; i < 10; i++ {
		if i != n {
			fmt.Println(i)
		}
	}
}
