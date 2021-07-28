package main

type InterfaceA interface {
	AAA(int)
	BBB(string)
}

// interfaceA.go 와 main.go 가 interface로 관계를 가지므로 출력값을 가짐
// String -> 퍼블릭함수를 참조
