package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 1.입력받을 소지금액을 담는 변수 생성.
	var cost int

	// 2.소지금을 유저로 부터 입력받는 함수 호출
	cost = input_user_Cost()
	// 무한반복문 돌림.
	for {
		// 3. 메뉴 출력
		printMenu()
		// 4. 메뉴를 번호로 받는 함수를 호출 후 변수에 담음
		no := inputMenuNumber()
		// 5. 번호를 담은 변수를 swith문을 통해 소지금을 감소시키는 함수 호출 후 소지금액 변수에 다시 담음.
		cost = switch_To_Cost(no, cost)
		// 6. 소지금액이 0 이하면 무한반복문 탈출하고, 그게 아니라면 남은 소지금 출력
		if cost <= 0 {
			fmt.Println("소지금액이 부족합니다.")
			break
		} else {
			fmt.Printf("남은 소지금액은 %d원입니다.", cost)
		}
	}
}

func input_user_Cost() int {
	var cost int
	for {
		fmt.Println("소지금은 1,000원 이상  10,000원 이하의 소지금을 입력해주세요.")
		fmt.Println("소지금을 입력하세요 (1,000원 ~ 10,000원) : ")

		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		n1, _ := strconv.Atoi(line)

		if n1 < 1000 || n1 >= 10000 {
			fmt.Println("1,000원 이상  10,000원 이하의 소지금을 입력해주세요.")
			continue

		} else {
			cost = n1
			break
		}
	}
	return cost
}

func printMenu() {
	fmt.Println("---------------------------------------------------------------------------------------------")
	fmt.Printf("\t\t\t\t\t\t\t\t음료수 메뉴\n")
	fmt.Printf("   1. 비락식혜 500원    2. 솔의 눈 1000원   3. 실론티 700원   4 이상. 나머지음료 1000원 \n")
	fmt.Println("---------------------------------------------------------------------------------------------")
}

func inputMenuNumber() int {
	fmt.Println("음료 번호를 입력하세요 : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	no, _ := strconv.Atoi(line)

	return no
}

func switch_To_Cost(no, cost int) int {
	switch no {
	case 1:
		fmt.Println("비락식혜가 나왔습니다.")
		cost = cost - 500

	case 2:
		fmt.Println("솔의 눈이 나왔습니다.")
		cost = cost - 1000

	case 3:
		fmt.Println("실론티가 나왔습니다.")
		cost = cost - 700

	default:
		fmt.Println("나머지 음료가 나왔습니다.")
		cost = cost - 1000
	}
	return cost
}
