package main

import (
	"fmt"
	"sort"
)

func main() {
	ary := make([]int, 3)
	fmt.Print("3つの値を入力: ")
	_, _ = fmt.Scan(&ary[0], &ary[1], &ary[2])
	sort.Slice(ary, func(i, j int) bool {
		return ary[i] < ary[j]
	})
	fmt.Println(ary[1])
}
