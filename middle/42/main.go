package main

import "fmt"

func main() {
	ary := make([]int, 3)
	for i := range ary {
		fmt.Printf("input number %d: ", i+1)
		_, _ = fmt.Scan(&ary[i])
	}

	if ary[0] <= ary[1] && ary[1] <= ary[2] {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
