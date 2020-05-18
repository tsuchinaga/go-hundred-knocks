package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("input 1st number: ")
	_, _ = fmt.Scan(&n)
	fmt.Print("input 2nd number: ")
	_, _ = fmt.Scan(&m)
	fmt.Printf("和: %d\n", n+m)
	fmt.Printf("差: %d\n", n-m)
	fmt.Printf("積: %d\n", n*m)
	fmt.Printf("商: %d, 余り: %d\n", n/m, n%m)
}
