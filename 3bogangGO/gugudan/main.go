package main

import "fmt"

func main() {
	for i := 2; i < 10; i++ {
		if (i%2) == 0 && ((i%4 == 0) || (i%2 == 0)) {
			fmt.Printf("----%d 단----\n", i)
		} else {
			continue
		}
		for j := 2; j < 10; j++ {
			if (j%2 == 0) || (j%4 == 0) {
				fmt.Printf("%d X %d = %d \n", i, j, i*j)
			}
		}
	}
}
