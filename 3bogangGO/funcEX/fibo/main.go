package main // 피보나치 수열,  주석 해제 후 작동 시 주의 무한 반복함

import "fmt"

func main() {
	rst := fibbonachi(10)
	fmt.Println(rst)
}

func fibbonachi(x int) int {
	if x == 0 {
		return 1
	} else if x == 1 {
		return 1
	}

	// fmt.Println("1 : ", fibbonachi(x-1))
	// fmt.Println("2 : ", fibbonachi(x-2))
	// fmt.Println("3 : ", fibbonachi(x-1)+fibbonachi(x-2))
	// fmt.Println("4 : ", fibbonachi(x))
	return fibbonachi(x-1) + fibbonachi(x-2)
}

// f(0) = 1
// f(1) = 1
// f(2) = (2-1) = 1 + (2-2) = 1 = 2
// f(3) = (3-1) = 2 + (3-2) = 1 = 3
// f(4) = (4-1) = 3 + (4-2) = 2 = 5
// f(5) = (5-1) = 5 + (5-2) = 3 = 8
