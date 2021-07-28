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
	// 1. 시도 횟수 변수 + 점수 변수를 생성
	var try int
	var score int

	try = 0
	score = 0

	print_to_Dest()
	// 2. 무한반복문으로 돌리기
	for {
		// 3. 랜덤 숫자 생성 함수 호충
		no1, no2 := randNumvers()
		// 4. 함수를 변수로 받아 랜덤으로 식을 내는 코드를 짜야됨.
		fmt.Printf("%d * %d의 값은  \n", no1, no2)
		// 5. 유저가 답을 입력하는 함수를 호출
		res := no1 * no2
		answer := inputUserAnswer()
		// 6. 유저가 입력한 변수와 랜덤식의 답을 체크하는 함수를 호출
		// 7. 맞출 때 마다 점수가 1점씩 상승, 틀리면 1점씩 하락하는 함수 호출
		// 8. 답변을 말할 떄 마다 시도횟수 1씩 증가하는 함수
		if check_to_Answer(answer, res) {
			fmt.Println("정답닙니다!")
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
		// 9. 5점을 얻었을 때 게임을 이기는 코드를 작성.
		if score == 5 {
			fmt.Println("축하합니다! 5점을 얻으셨습니다.")
			break
		}
	}
	fmt.Printf("총 시도횟수는 %d번 입니다.", try)
}

func randNumvers() (int, int) {
	//시드는 언제나 바뀌게
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10)
	y := rand.Intn(10)

	return x, y
}

func inputUserAnswer() int {
	fmt.Println("정답은 ? : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	answer, _ := strconv.Atoi(line)

	return answer
}

func check_to_Answer(x, y int) bool {
	if x == y {
		fmt.Println("정답입니다.")
		return true
	} else {
		return false
	}
}

func print_to_Dest() {
	fmt.Println("구구단 게임을 시작합니다.")
	fmt.Println("랜덤 숫자를 2개 컴퓨터기 정한 뒤 유저가 답을 입력하는 방식입니다.")
	fmt.Println("정답을 맞추면 1점을 얻고, 정답이 틀리면 1점이 감소가 됩니다.")
	fmt.Println("5점을 얻었습니다.")
}
