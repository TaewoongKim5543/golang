package main // 함수를 이용한 계산기 만들기

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sum(10, 5))
	fmt.Println(subtract(10, 5))
	fmt.Println(division(10, 5))
	fmt.Println(multiply(10, 5))

	//계산기 숫자를 입력 받기
	num01, num02 := inputNumberSec()
	//연산자 입력 받기
	Operator := inputOperator()
	//결과 연산하기
	result := calculatorNum(num01, num02, Operator)
	//결과 값 출력하기
	printResult(num01, num02, Operator, result)

}

func sum(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func division(x, y int) int {
	return x / y
}

func multiply(x, y int) int {
	return x * y
}

// 숫자 입력하라는 메세지 출력함수
func inputNumberMsg() {
	fmt.Println("숫자를 입력하세요. : ")
}

//연산자 입력하라는 메세지 출력함수
func inputOperatorMsg() {
	fmt.Println("연산자를 입력하세요. : ")
}

// 계산기~
// 1. 계산할 숫자 2개를 입력받아 2개의 숫자를 반환하는 함수
func inputNumberSec() (int, int) { // Reader 객체를 사용해서!
	// 문구 출력
	inputNumberMsg()
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	num01, _ := strconv.Atoi(line)

	// 문구 출력
	inputNumberMsg()
	reader02 := bufio.NewReader(os.Stdin)
	line02, _ := reader02.ReadString('\n')
	line02 = strings.TrimSpace(line02)
	num02, _ := strconv.Atoi(line02)

	return num01, num02 // 숫자 두개 반환
}

// 2. 연산할 연산자를 입력받아 연산자를 받은 문자열 변수를 반환하는 함수
func inputOperator() string { // Reader 객체 가능 , Scanf 가능
	// 문구 출력
	inputOperatorMsg()
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(line)

	return operator // 연산자 반환
}

// 3. 입력받은 숫자를 반환한 가변인자 2개와 연산자를 반환한 가변인자 1개를 받아
// 조건문을 통해 연산을 하여 연산 결과를 반환하는 함수
func calculatorNum(x, y int, o string) int {
	var result int

	switch o {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		result = x / y
	default:
		result = 0 // 함수가 종료될때 까지 지연, 반환 후 끝나면 출력
		defer fmt.Println("연산자가 아닙니다.")
	}
	return result
}

// 4. 연산한 결과를 형식에 맞게 출력하는 함수
func printResult(x, y int, o string, res int) {
	fmt.Printf("%d %s %d의 값은 = %d이다.", x, o, y, res)
}
