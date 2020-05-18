package main

import "fmt"

func main() {
	ary := [...]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2}

	var n, m int
	fmt.Print("input index1: ")
	_, _ = fmt.Scan(&n)
	fmt.Print("input index2: ")
	_, _ = fmt.Scan(&m)

	fmt.Printf("%d * %d = %d\n", ary[n], ary[m], ary[n]*ary[m])
}
