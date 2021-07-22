package main // 재귀함수와 반복문 비교

import "fmt"

//---------------------------------------------------------------------//
// func main() { //재귀방식
// 	rst := sum(10, 0)

// 	fmt.Println(rst)
// }

// func sum(x int, s int) int {
// 	if x == 0 {
// 		return s
// 	}

// 	s += x
// 	return sum(x-1, s)
// }
//---------------------------------------------------------------------//
func main() { //반복문 방식
	s := 0
	for i := 1; i <= 10; i++ {
		s += i
	}
	fmt.Println(s)
}
