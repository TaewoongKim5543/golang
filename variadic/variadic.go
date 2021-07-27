package main //가변인수함수 연습예제

import (
	"fmt"
	"os"
)
// ... 요거를 인수 타입앞에 붙이면 해당 타입 인수를 여러개 받는 가변인수임을 표시함.
func varFunc(input ...string) { //가변인수로 string을 받는 함수 구현. 인수 input은 함수 내부에서 슬라이스로 취급/ ...=> 팩 연산자
	fmt.Println(input) 
}

//-------------------------------------------------------------------------------
func oneByOne(message string, s ...int) int { //스트링하나 받고 정수는 가변인수로 받는 함수. S는 슬라이스
	fmt.Println(message)
	sum := 0
	for i, a := range s {
		fmt.Println(i, a)
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

//-------------------------------------------------------------------------------
func main() { // 앞의 가변인수 함수 2개 이용
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByOne("Adding numbers...", 1, 2, 3, 4, 5, -1, 10)
	fmt.Println("Sum:", sum)
	s := []int{1, 2, 3}
	sum = oneByOne("Adding numbers...", s...) // 슬라이스 사용, 함수 내부 슬라이스 변경 시 함수 종료시에도 결과 남음.
	fmt.Println(s)
}
