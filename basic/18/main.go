package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	ary := make([]int, 10)
	for i := range ary {
		ary[i] = n
	}
	for _, m := range ary {
		fmt.Println(m)
	}
}
