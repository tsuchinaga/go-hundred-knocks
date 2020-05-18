package main

import "fmt"

func main() {
	ary := [...]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2}

	var n int
	fmt.Print("input index: ")
	_, _ = fmt.Scan(&n)

	fmt.Printf("array[%d] = %d\n", n, ary[n])
}
