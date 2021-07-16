package main

import "fmt"

func main() {

	// 변수 선언
	a := 10
	b := 12

	// 산술 연산
	c := a + b
	d := a - b
	e := a / b
	f := a * b

	// 출력
	fmt.Println(c, d, e, f)
	fmt.Println()
	fmt.Println()
	// 논리 연산
	fmt.Println("a > b", (a > b))
	fmt.Println("a < b", (a < b))
	fmt.Println("a == b", (a == b))
	fmt.Println("a != b", (a != b))
	fmt.Println("(a > b) && (a < b)", (a > b) && (a < b))
	fmt.Println("(a > b) || (a < b)", (a > b) || (a < b))
	fmt.Println()
	// 비트연산(단항 연산)
	fmt.Println("a & b", a&b)
	fmt.Println("a | b", a|b)
	fmt.Println("a ^ b", a^b)
	fmt.Println("a &^ b", a&^b)
	fmt.Println()
	fmt.Println()
	//  쉬프트 연산

	a = 1024
	b = 3

	fmt.Println("a << b", a<<b)
	fmt.Println("a >> b", a>>b) //a 12칸 쉬프트 연산 해버려서 0됨.

}
