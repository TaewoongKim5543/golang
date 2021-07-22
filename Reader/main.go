package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 입력받을 소지 금액 값 담는 변수 생성
	var cost int
	for {
		fmt.Println("소지하고 있는 금액은 10000 이상 20000이하만 입력하세요!")
		fmt.Println("소지하고 있는 금액을 입력하세요 (10000 ~ 20000):  ")
		// bufio = 버퍼라는 공간에 데이터를 읽는 객체 생성
		// (운영체제 기본 입력 객체 (Stdin))
		reader := bufio.NewReader(os.Stdin)
		// 데이터를 읽는 객체 reader가 엔터가 쳐진 문자열 공간을
		// 버퍼 reader 객체가 읽어 line 이라는 변수에 대입.
		line, _ := reader.ReadString('\n')
		// 문자열에서 공백이나 그 외 쓰레기 값을 제외 시켜버림.
		// TrimSpace() : strings 패키지에 있는 메서드임
		line = strings.TrimSpace(line)
		// strconv : 문자열 컨버터 /  line 변수 안에 있는 문자열 값을
		// Atoi() : INT 타입으로 변환 시켜줌 -> a = int(a) 랑 똑같은 역할 .
		n1, _ := strconv.Atoi(line)
		// 유저가 입력한 금액이 10000 미만 20000 초과 일 경우
		if n1 < 10000 || n1 > 20000 {
			fmt.Println("10000 이상 20000 이하인 금액을 입력하세요.")
			continue
			// 그의 반대 일 경우는
		} else {
			// 소지금액 변수 안에 입력받은 숫자 대입.
			cost = n1
			break
		}
	}
	for {
		// 메뉴 설명
		fmt.Println("-------------------------------------------------------------------------------------------------------")
		fmt.Printf("\t\t\t\t\t\t\t\t음료수 메뉴\n")
		fmt.Printf("   1. 콜라 500원    2. 사이다 700원   3. 생수 600원   4. 게토레이 1200원   5 이상. 무한 나머지음료 1000원 \n")
		fmt.Println("------------------------------------------------------------------------------------------------------")
		fmt.Println("마실 음료수를 고르세요.")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		no, _ := strconv.Atoi(line)
		// 메뉴 골랐을 때 각 음료수마다 가격에 따라 소지금액 감산
		switch no {
		case 1:
			fmt.Println("콜라가 나왔습니다.")
			cost = cost - 500
		case 2:
			fmt.Println("사이다가 나왔습니다.")
			cost = cost - 700
		case 3:
			fmt.Println("생수가 나왔습니다.")
			cost = cost - 600
		case 4:
			fmt.Println("게토레이가 나왔습니다.")
			cost = cost - 1200
		default:
			fmt.Println("그외 음료수가 나왔습니다.")
			cost = cost - 1000
		}
		// 소지금액이 0원 이거나 0 미만일 경우 무한 반복문 탈출
		if cost <= 0 {
			fmt.Println("소지 금액이 부족합니다.")
			break
		} else {
			// 소비하고 남은 돈 출력!
			fmt.Printf("남은 소지 금액은 %d입니다.\n", cost)
		}
	}
}
