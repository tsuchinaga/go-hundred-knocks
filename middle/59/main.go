package main

import "fmt"

func main() {
	fmt.Println("1つめの行列")
	ary1 := make([][]int, 3)
	for i := range ary1 {
		ary := make([]int, 3)
		_, _ = fmt.Scan(&ary[0], &ary[1], &ary[2])
		ary1[i] = ary
	}

	fmt.Println("2つめの行列")
	ary2 := make([][]int, 3)
	for i := range ary1 {
		ary := make([]int, 3)
		_, _ = fmt.Scan(&ary[0], &ary[1], &ary[2])
		ary2[i] = ary
	}

	fmt.Println("和")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\t%d\t%d\n", ary1[i][0]+ary2[i][0], ary1[i][1]+ary2[i][1], ary1[i][2]+ary2[i][2])
	}
}
