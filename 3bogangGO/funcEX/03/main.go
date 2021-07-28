package main // 재귀 함수 실습

import (
	"fmt"
)

func main() {
	f1(10)
	// f2(5)
}

func f1(x int) {
	if x == 0 {
		return
	}
	fmt.Printf("f1(%d) before call f1(%d)\n", x, x-1)
	f1(x - 1)
	fmt.Printf("f1(%d) after call f1(%d)\n", x, x-1)
}

// func f2(y int) {
// 	if y == 0 {
// 		return
// 	}
// 	f2(y - 1)
// 	fmt.Printf("f2(%d) before call f2(%d)\n", y, y-1)
// 	fmt.Printf("f2(%d) after call f2(%d)\n", y, y-1)
// }
