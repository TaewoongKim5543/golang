package main

import "fmt"

func main() {
	var money int
	var menu string

	for {
		fmt.Printf("소지하고 있는 금액은 10000이상 20000이하만 입력하세요! \n소지하고 있는 금액을 입력하세요(10000 ~ 20000)  \n-----------------------------------------------------\n")
		_, err := fmt.Scanf("%d\n", &money)
		if err != nil {
			fmt.Println("잘못 입력했습니다.")
		} else {
			if money < 10000 || money > 20000 {
				continue
			} else {
				break
			}
		}
	}
	for {
		for {
			fmt.Print("-------------------------------------------------------------------------------------------------------------\n                                                 음료수 메뉴 \n 1. 콜라 500원 2. 사이다 700원  3. 생수 600원  4. 게토레이 1200원  5. 이상 무한 나머지 음료 1000원 \n-------------------------------------------------------------------------------------------------------------\n ")
			_, err := fmt.Scanf("%s\n", &menu)
			if err != nil {
				fmt.Println("잘못 입력했습니다.")
			} else {
				break
			}
		}
		switch menu {
		case "1":
			money = money - 500
			fmt.Printf("콜라가 나왔습니다. \n남은 소지금액은 %d입니다.\n", money)
		case "3":
			money = money - 600
			fmt.Printf("생수가 나왔습니다. \n남은 소지금액은 %d입니다.\n", money)
		case "4":
			money = money - 1200
			fmt.Printf("게토레이가 나왔습니다. \n남은 소지금액은 %d입니다.\n", money)
		case "2":
			money = money - 700
			fmt.Printf("사이다가 나왔습니다. \n남은 소지금액은 %d입니다.\n", money)
		default:
			money = money - 1000
			fmt.Printf("남은 소지금액은 %d입니다.\n", money)
		}
		if money <= 0 {
			fmt.Println("소지금액이 부족합니다.")
			break
		}
	}
}
