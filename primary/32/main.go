package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%5 == 0 {
			fmt.Println("bar")
		} else {
			fmt.Println(i)
		}
	}
}
