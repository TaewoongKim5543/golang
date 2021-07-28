package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 2; i < 10; i++ {
		fmt.Printf("%d 제곱단\n", i)
		for j := 2; j < 10; j++ {
			fmt.Printf("%d의 %d 제곱 = %d\n", i, j, int(math.Pow(float64(i), float64(j))))
			// fmt.Print(math.Pow10(i))
		}
		fmt.Println("-------------")
	}
}
