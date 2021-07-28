package main

import "fmt"

func main() {
	fmt.Println(fi(9))
}

func fi(x int) int {
	if x == 0 {
		return 1
	} else if x == 1 {
		return 1
	}

	return fi(x-1) + fi(x-2)
}
