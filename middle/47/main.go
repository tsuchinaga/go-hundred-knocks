package main

import "fmt"

func main() {
	ary := make([]int, 2)
	for i := range ary {
		fmt.Printf("input %s: ", string('a'+i))
		_, _ = fmt.Scan(&ary[i])
	}
	ary[0], ary[1] = ary[1], ary[0]
	fmt.Printf("a = %d, b = %d\n", ary[0], ary[1])
}
