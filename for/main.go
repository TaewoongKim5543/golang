package main

import "fmt"

func main() {
	//기본 반복문

	for i := 0; i < 10; i++ {
		fmt.Printf("현재 i 값은 : %d\n", i)
	}
	//무한 반복문
	a := 0
	for {
		if a == 12 {
			fmt.Println("탈출한다.")
			break
		}
		a++
	}
	//조건 반복문
	j := 0
	for j < 20 {
		fmt.Printf("j의 값은 %d이다.\n", j)
		j++
	}
}
