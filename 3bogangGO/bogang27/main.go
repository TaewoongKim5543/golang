package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var try int
	var score int

	print_to_Dest()

	for {
		no1, no2 := randNumbers()
		fmt.Printf("%d * %d의 값은 ? \n", no1, no2)
		res := no1 * no2
		answer := inputUserAnswer()

		if check_to_Answer(answer, res) {
			fmt.Println("정답입니다.")
			fmt.Println("1점을 얻었습니다.")
			score++
			try++
			fmt.Printf("현재 점수 : %d\n", score)
		} else {
			fmt.Println("틀렷습니다.")
			fmt.Println("1점이 감소됩니다.")
			score--
			try++
			fmt.Printf("현재 점수 : %d\n", score)
		}
		if score == 5 {
			fmt.Println("축하합니다! 5점을 얻으셨습니다.")
			break
		}
	}
	fmt.Printf("총 시도횟수는 %d번입니다.", try)
}

func randNumbers() (int, int) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)

	return x, y
}

func inputUserAnswer() int {
	fmt.Println("정답 : ")
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
	// Atoi() : INT 타입으로 변환 시켜줌
	answer, _ := strconv.Atoi(line)
	// 적은 답이 컴퓨터에 저장된 정답과 같으면

	return answer
}

func check_to_Answer(x, y int) bool {
	if x == y {
		return true
	} else {
		return false
	}
}

func print_to_Dest() {
	fmt.Println("구구단 게임을 시작합니다.")
	fmt.Println("랜덤 숫자를 2개를 컴퓨터가 정한 뒤 유저가 답을 입력하는 방식입니다.")
	fmt.Println("정답을 맞추면 1점을 얻고, 정답이 틀리면 1점이 감소가 됩니다.")
	fmt.Println("5점을 얻었을때, 게임이 종료됩니다.")
	fmt.Println("그럼, 구구단 게임을 시작하겠습니다!")
}
