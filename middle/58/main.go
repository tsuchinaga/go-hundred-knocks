package main

import (
	"fmt"
	"strings"
)

func main() {
	ary := make([]int, 5)
	for i := range ary {
		fmt.Printf("input data[%d]: ", i)
		_, _ = fmt.Scan(&ary[i])
	}

	for _, n := range ary {
		s := make([]string, 0)
		for i := 5; i <= n; i += 5 {
			s = append(s, "*****")
		}
		if n%5 > 0 {
			s = append(s, strings.Repeat("*", n%5))
		}
		fmt.Printf("%d\t:%s\n", n, strings.Join(s, " "))
	}
}
