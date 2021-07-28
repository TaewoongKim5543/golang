// package main // 배열 예제 연습

// import "fmt"

// func main() {
// 	var A [10]int

// 	for i := 0; i < 10; i++ { // 10은 len(A) 로 대체가능 A의 길이 는 10이니까~
// 		A[i] = i * i
// 	}
// 	fmt.Println(A)
// }
//////////////////////////////////////////////////////////////////////////////
// package main // 문자열 예제 연습~

// import (
// 	"fmt"
// )

// func main() {
// 	s := "Hello World"

// 	for i := 0; i < len(s); i++ {
// 		// fmt.Print(s[i], ", ")
// 		fmt.Print(string(s[i]), ", ")
// 	}
// }
//////////////////////////////////////////////////////////////////////////////
package main // 예제 연습2~

import (
	"fmt"
)

func main() {
	s := "Hello 와아아알드"
	s2 := []rune(s)

	for i := 0; i < len(s2); i++ {
		fmt.Print(s2[i], ", ")

	}
}
