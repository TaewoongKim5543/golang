package main

import (
	"fmt"
)

type StructA struct {
	val string
}

func (s *StructA) String() string { // String -> 퍼블릭
	return "Val:" + s.val
}

func main() {
	a := &StructA{val: "AAA"}
	fmt.Println(a)
}
