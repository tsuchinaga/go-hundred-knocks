package main

import "fmt"

func main() {
	var n int
	fmt.Print("input number: ")
	_, _ = fmt.Scan(&n)

	if (-10 <= n && n < 0) || 10 <= n {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
