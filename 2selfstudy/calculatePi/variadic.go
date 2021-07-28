package main

import (
	"fmt"
	"os"
)

func varFunc(input ...string) {  // 가변인수로 스트링을 받는 함수
	fmt.Println(input) //input은 슬라이스 취급 / ... 연산자 -> 팩연산자
}

func oneByone(message string, s ...int) int { // 스트링하나받고 정수를 가변인수로 받는 함수. s는 슬라이스
	fmt.Println(message)
	sum := 0
	for i, a := range s {
		fmt.Println(i, a)
		sum = sum + a
	}
	s[0] = -1000
	return sum
}

func main() {
	arguments := os.Args
	varFunc(arguments...)
	sum := oneByone("Adding numbers...", 1, 2, 3, 4, 5, -1, 10)
	fmt.Println("sum:", sum)
	s := []int{1, 2, 3}
	sum = oneByone("Adding numbers...", s...)
	fmt.Println(s)
}