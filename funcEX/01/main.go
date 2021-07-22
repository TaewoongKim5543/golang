package main // 함수 기초

import "fmt"

func add(a int, b int) int {
	return (a + b)
}

func main() {

	c := add(4, 5)

	fmt.Println(c)
}
