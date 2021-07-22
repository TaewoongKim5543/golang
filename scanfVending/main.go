package main

import "fmt"

func main() {
	// 1. 주석 입력 - 돈과 메뉴를 입력받을 변수를 생성한다.
	var cost int
	var menu int

	// 2. 소지금을 입력할 문자 출력과 입력함수를 이용해 처음 소지금을 입력 한다.
	for {
		fmt.Println("소지금은 10,000 이상 20,000 이하로 입력해야 한다.")
		fmt.Println("소지금을 입력하세요 (10,000원 ~ 20,0000원) : ")

		// *추가 - 정수가 아닌 다른 타입이 들어왔을 때 다시 입력해야한다.
		_, err := fmt.Scanf("%d\n", &cost) // 빈 변수 설정 이유 뒤의 cost를 받으려고
		if err != nil {
			fmt.Println("잘 못 입력하였습니다.")
		} else {
			// 10,000 이상 20,000 이하 조건을 충족한 숫자를 입력 했을 경우
			if (cost >= 10000) && (cost <= 20000) {
				break
			} else {
				//반대일 경우 계속
				continue
			}
		}
	}

	// 3. 무한 반복문을 통해 메뉴를 출력하고 사용자가 메뉴를 입력한다.
	for {
		for {
			fmt.Println("-----------------------------------------------------------------------------------------")
			fmt.Printf("\t\t\t음료메뉴\n")
			fmt.Println("1. 콜라 600원 2. 사이다 700원 3. 생수 500원 4. 게토레이 1500원 5. 그외 나머지 음료 1000원 ")
			fmt.Println("-----------------------------------------------------------------------------------------")

			_, err01 := fmt.Scanf("%d\n", &menu)
			// *추가 - 정수가 아닌 다른 타입이 들어왔을 때 다시 입력해야한다.
			if err01 != nil {
				fmt.Println("잘 못 입력했습니다.")
			} else {
				break
			}
		}

		switch menu {

		case 1: // 4. 메뉴를 입력했을 때 숫자로 입력 받으면 소지금에서 그 메뉴의 가격에 따라
			fmt.Println("콜라가 나왔습니다.")
			fmt.Println("600원이 감소하였습니다.")
			cost = cost - 600 // 5. 소지금을 감한다.
		case 2:
			fmt.Println("사이다가 나왔습니다.")
			fmt.Println("700원이 감소하였습니다.")
			cost = cost - 700
		case 3:
			fmt.Println("생수가 나왔습니다.")
			fmt.Println("500원이 감소하였습니다.")
			cost = cost - 500
		case 5:
			fmt.Println("게토레이가 나왔습니다.")
			fmt.Println("1500원이 감소하였습니다.")
			cost = cost - 1500
		default:
			fmt.Println("그 외 음료수가 나왔습니다.")
			fmt.Println("1200원이 감소하였습니다.")
			cost = cost - 1200
		}
		//8. 남은 소지금액 표기 추가
		fmt.Printf("남은 소지금액은 %d이다.", cost)

		// 6. 연속적으로 반복하다 소지금이 부족해지면 or 소지금액이 0 이하일 경우.
		// 7. 소지 금액 부족 경고문을 출력 후 무한 반복문 탈출.
		if cost <= 0 {
			fmt.Println("소지 금액이 부족합니다.")
			break
		}
	}
}
