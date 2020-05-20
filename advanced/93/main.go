package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var start, end int
	fmt.Print("start: ")
	_, _ = fmt.Scan(&start)
	fmt.Print("end: ")
	_, _ = fmt.Scan(&end)
	for i := start; i <= end; i++ {
		head := ""
		if i%3 == 0 || strings.Contains(strconv.Itoa(i), "3") {
			head = "aho "
		}
		fmt.Printf("%s%d\n", head, i)
	}
}
