package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for i := 1; i <= 50; i++ {
		head := ""
		if i%3 == 0 || strings.Contains(strconv.Itoa(i), "3") {
			head = "aho "
		}
		fmt.Printf("%s%d\n", head, i)
	}
}
