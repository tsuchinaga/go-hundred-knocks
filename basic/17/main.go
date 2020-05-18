package main

import "fmt"

func main() {
	ary := make([]int, 10)
	for i := range ary {
		ary[i] = i
	}
	for _, n := range ary {
		fmt.Println(n)
	}
}
