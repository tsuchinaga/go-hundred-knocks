package main

import "fmt"

func main() {
	ary := [...]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2}
	for i := 0; i < len(ary)-1; i++ {
		fmt.Println(ary[i] - ary[i+1])
	}
}
