package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("何円? ")
	_, _ = fmt.Scan(&n)
	fmt.Print("1ドルは何円? ")
	_, _ = fmt.Scan(&m)
	fmt.Printf("%d円は%dドル%dセント\n", n, n/m, n%m*100/m)
}
