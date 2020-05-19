package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// data/small.data, data/middle.data, data/large.dataをリダイレクトで渡す
// small.dataの回答では最小が128となっているが、データでは最小が10なので10を正答としていいと思う
func main() {
	max, min := math.MinInt64, math.MaxInt64
	bs := bufio.NewScanner(os.Stdin)
	bs.Split(bufio.ScanLines)
	for bs.Scan() {
		n, _ := strconv.Atoi(bs.Text())
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Printf("最小値 = %d, 最大値 = %d\n", min, max)
}
