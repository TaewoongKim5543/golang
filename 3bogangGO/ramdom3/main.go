package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// float64() 64비트 실수 범위 사이의 실수를 아무거나 랜덤 실수 생성.
	// 100~ 200 사이의 랜덤숫자를 뽑는다 (* 100은 첫번째 범위 (+100)은 100~200사이 길이를 더해주기 )
	for {
		none := false
		a := int(rand.New(rand.NewSource(time.Now().UnixNano())).Float64()*100) + 100
		for i := 2; i < 10; i++ {
			if a%i == 0 {
				none = false
				fmt.Printf("%d는 %d의 배수입니다. \n", a, i)
				break
			} else {
				none = true
			}
		}
		if none {
			fmt.Printf("%d는 2~9 사이의 어느 배수도 아니다 \n", a)
		} else {
			break
		}
	}
}
