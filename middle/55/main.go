package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print("とんで")
		}
		for j := 0; j < 3; j++ {
			fmt.Print("まわって")
		}
		fmt.Println("まわる")
	}
}
