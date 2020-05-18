package main

import "fmt"

func main() {
	ary := [...]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2}
	prev := 0
	for i := 0; i < len(ary); i++ {
		fmt.Println(ary[prev])
		prev = ary[prev]
	}
}
