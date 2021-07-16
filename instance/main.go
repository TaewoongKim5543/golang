// package main //instance 방식 코딩 기본! #

// import (
// 	"fmt"
// )

// type Student struct {
// 	name  string
// 	age   int
// 	grade int
// }

// func main() {
// 	a := Student{"김태웅", 29, 10}
// 	//  #1. none pointer
// 	//  b := a >> a 값 직접 변경
// 	//  b.age = 20
// 	b := &a // a값 카피 후 변경

// 	fmt.Println(a)
// 	fmt.Println(b)
// }
///////////////////////////////////////////////////////////////////////////////////////
// #2 함수를 통한 갱신
package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func SetName(t *Student, newName string) { // 스튜던트 주소로가라
	t.name = newName
}

func main() {
	a := Student{"김태웅", 29, 10}

	fmt.Println(a)
	SetName(&a, "감태웅") //a로 가서 김태웅을 감태웅으로 바꿔라
	fmt.Println(a)
}
