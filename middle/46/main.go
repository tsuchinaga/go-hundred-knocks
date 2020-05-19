package main

import "fmt"

func main() {
	var n int
	fmt.Print("人数 ")
	_, _ = fmt.Scan(&n)

	t := 600
	if n >= 20 {
		t = 500
	} else if n >= 5 {
		t = 550
	}
	fmt.Printf("料金: %d\n", n*t)
}
