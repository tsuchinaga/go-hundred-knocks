package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	avg := make([]int, 3)
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		var a, b, c int
		_, _ = fmt.Scan(&a, &b, &c)
		avg[0] += a
		avg[1] += b
		avg[2] += c
		sum[i] = a + b + c
	}

	fmt.Printf("平均点 英語:%d, 数学:%d, 国語:%d\n", avg[0]/n, avg[1]/n, avg[2]/n)
	fmt.Println("個人合計点")
	for i, s := range sum {
		fmt.Printf("%d: %d\n", i, s)
	}
}
