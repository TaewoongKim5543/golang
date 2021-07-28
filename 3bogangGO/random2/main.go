package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(200)
	b := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	var count int

	for i := 0; i <= a; i++ {
		for j := 0; j <= a; j++ {
			if (i + j) == b {
				fmt.Printf("(%d,%d)\n", i, j)
				count++
			}
		}
	}

	fmt.Printf("첫 번째 랜덤 숫자는 %d다\n", a)
	fmt.Printf("두 번째 랜덤 숫자는 %d다\n", b)
	fmt.Printf("총 출력한 갯수는 %d개이다\n", count)

}
