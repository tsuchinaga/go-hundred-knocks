package main

import "fmt"

func main() {
	ary := make([]int, 5)
	for i := range ary {
		var n int
		fmt.Print("input number: ")
		_, _ = fmt.Scan(&n)
		ary[i] = n
	}
	for _, m := range ary {
		fmt.Println(m)
	}
}
