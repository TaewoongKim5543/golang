package main

import "fmt"

func main() {
	var a int
	_, err := fmt.Scanf("%d", &a)

	if err != nil {
		fmt.Println("잘못 입력하였습니다.")
	}

	if a > 10 {
		fmt.Println("a는 10보다 크다!")
	} else {
		fmt.Println("a는 10보다 작다")
	}
}
