package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		s := ""
		if i%3 == 0 {
			s += "foo"
		}
		if i%5 == 0 {
			s += "bar"
		}
		if s == "" {
			s = strconv.Itoa(i)
		}
		fmt.Println(s)
	}
}
