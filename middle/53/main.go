package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	pn := 2
	pns := make([]string, 0)
	for {
		if n == 1 {
			break
		}

		if n%pn == 0 {
			n /= pn
			pns = append(pns, strconv.Itoa(pn))
		} else {
			pn++
		}
	}
	fmt.Println(strings.Join(pns, " "))
}
