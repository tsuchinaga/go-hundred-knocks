package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j != 1 {
				fmt.Print("\t")
			}
			fmt.Print(i * j)
		}
		fmt.Println()
	}
}
