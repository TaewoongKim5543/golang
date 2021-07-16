// package main

// import "fmt"

// func main() {
// 	var a int
// 	var b int

// 	var op string

// 	fmt.Print("첫 번째 숫자를 입력하세요 : \n")
// 	_, err := fmt.Scanf("%d\n", &a)

// 	fmt.Print("두 번째 숫자를 입력하세요 : \n")
// 	_, err01 := fmt.Scanf("%d\n", &b)
// 	if err != nil {
// 		fmt.Println("잘 못 입력하였습니다. ")
// 	}

// 	if err01 != nil {
// 		fmt.Println("잘 못 입력하였습니다. ")
// 	}

// 	fmt.Print("연산할 연산자를 입력하세요.\n")

// 	_, err02 := fmt.Scanf("%s\n", &op)

// 	if err02 != nil {
// 		fmt.Println("연산자를 잘 못 입력하였습니다.")
// 	}

// 	// if op == "+" {
// 	// 	fmt.Println("a + b = ", a+b)
// 	// } else if op == "-" {
// 	// 	fmt.Println("a - b = ", a-b)
// 	// } else if op == "*" {
// 	// 	fmt.Println("a * b = ", a*b)
// 	// } else if op == "/" {
// 	// 	fmt.Println("a / b = ", a/b)
// 	// } else {
// 	// 	fmt.Println("연산자가 아닙니다.")
// 	// }

// 	switch op {
// 	case "+":
// 		fmt.Println("a +b", a + b)
// 	case
// 		fmt.Println("a - b = ", a-b)
// 	case
// 		fmt.Println("a * b = ", a*b)
// 	case
// 		fmt.Println("a / b = ", a/b)
// 	default
// 		fmt.Println("연산자가 아닙니다.")
// 	}
// }

package main

import "fmt"

func main() {
	var num01 int
	var num02 int
	var op string

	fmt.Println("첫 번째 숫자를 입력하시오 :")
	_, err01 := fmt.Scanf("%d\n", &num01)

	if err01 != nil {
		fmt.Println("잘 못 입력했습니다.")
	}

	fmt.Println("두 번째 숫자를 입력하시오 :")
	_, err02 := fmt.Scanf("%d\n", &num02)

	if err02 != nil {
		fmt.Println("잘 못 입력했습니다.")
	}

	fmt.Printf("연산할 연산자를 입력하세요(+, -, /, *)")

	_, err03 := fmt.Scanf("%s\n", &op)

	if err03 != nil {
		fmt.Println("잘 못 입력했습니다.")
	}

	// if op == "+"{
	// 	fmt.Printf("이 식의 값은 %d, num01 + num02")
	// }else if op == "-"{
	// 	fmt.Printf("이 식의 값은 %d, num01 - num02")
	// }else if op == "*"{
	// 	fmt.Printf("이 식의 값은 %d, num01 * num02")
	// }else if op == "/"{
	// 	fmt.Printf("이 식의 값은 %d, num01 / num02")

	switch op {
	case "+":
		fmt.Printf("이 식의 값은 %d", num01+num02)
	case "-":
		fmt.Printf("이 식의 값은 %d", num01-num02)
	case "/":
		fmt.Printf("이 식의 값은 %d", num01/num02)
	case "*":
		fmt.Printf("이 식의 값은 %d", num01*num02)
	default:
		fmt.Printf("잘 못 입력했습니다.")
	}
}
