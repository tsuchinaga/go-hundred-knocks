package main

import "fmt"

func main() {
	ary := make([]int, 3)
	for i := range ary {
		fmt.Printf("input %s: ", string('a'+i))
		_, _ = fmt.Scan(&ary[i])
	}

	d := ary[1]*ary[1] - 4*ary[0]*ary[2]
	if d > 0 {
		fmt.Println("2つの実数解")
	} else if d < 0 {
		fmt.Println("2つの虚数解")
	} else {
		fmt.Println("重解")
	}
}
