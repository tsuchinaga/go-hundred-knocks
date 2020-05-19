package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("2つの値をスペースで区切って入力: ")
	_, _ = fmt.Scan(&n, &m)

	for {
		if n%m == 0 {
			break
		}
		n, m = m, n%m
	}

	if m == 1 {
		fmt.Println("互いに素")
	} else {
		fmt.Println("互いに素でない")
	}
}
