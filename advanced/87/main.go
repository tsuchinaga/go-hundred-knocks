package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Print("誕生日をYYYYMMDDの形式で入力してください: ")
	_, _ = fmt.Scan(&s)

	for {
		var sum int32
		for _, c := range s {
			sum += c - '0'
		}
		s = strconv.Itoa(int(sum))

		// 1桁か、ゾロ目かのチェック
		// 最大9 x 8桁の72なので、2桁以上を考慮する必要はない
		if len(s) == 1 || s[0] == s[1] {
			break
		}
	}
	fmt.Printf("運命数は%s\n", s)
}
