package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// a := 5
	// for i := 0; i < a; i++ {
	// 	fmt.Println(i)
	// }
	// a = 0
	// for a < 5 {
	// 	fmt.Println(a)
	// 	a++
	// }

	a := 0
	for a < 5 {
		if a%2 == 0 && a != 0 {
			fmt.Println("a는 2의 배수입니다. : ", a)
		}
		fmt.Println(a)
		a++
	}
}
