package main

import "fmt"

func main() {
	ary := []int{1}
	for i := 0; i < 15; i++ {
		output(ary)

		ary = append([]int{0}, append(ary, 0)...)
		nAry := make([]int, len(ary)-1)
		for i := 0; i < len(ary)-1; i++ {
			nAry[i] = ary[i] + ary[i+1]
		}
		ary = nAry
	}
}

func output(ary []int) {
	for i, n := range ary {
		if i != 0 {
			fmt.Print(" ")
		}
		fmt.Print(n)
	}
	fmt.Println()
}
