package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
	for i := 0; i < 5; i++ {
		for j := 0; j <= 4-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ { // K라는 변수를 추가해 주어야 얘가 정상작동, j로 하면 이상하게 뜸.
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("---------------------------------------")
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ { // K라는 변수를 추가해 주어야 얘가 정상작동, j로 하면 이상하게 뜸.
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 6; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 5-(i*2); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
