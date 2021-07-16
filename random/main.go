package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 5; i++ {
		ran := rand.Intn(50)
		var sum int
		for j := 1; j < ran; j++ { // sum연산 ..
			sum = sum + j
		}
		fmt.Printf("%d의 1 ~ %d사이의 sum(%d) = %d\n", ran, ran, ran, sum)
	}
}
