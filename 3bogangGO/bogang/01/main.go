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
	// 컴퓨터에 저장된 정답 값 담는 변수 생성
	var res int
	// 시도 횟수 값 담은 변수 생성
	var try int
	// 점수 값 담는 변수 생성
	var grade int
	fmt.Println("구구단 게임에서 컴퓨터가 내는 문제를 맞춰 5점을")
	fmt.Println("얻으면 게임은 끝나게 됩니다. (맞출 때 마다 1점씩 증가)")
	fmt.Println("단, 정답이 틀리면 1점 감산됩니다.")
	fmt.Println("그럼, 구구단 게임을 시작하겠습니다!")
	// 랜덤 숫자 시드 현재 식 나노 시간 까지 쪼개어 고정!
	// (언제나 임의의 숫자가 발행되게 하는 메서드)
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("컴퓨터가 문제를 구하는 중입니다.")
		//랜덤 숫자 2개 0 ~ 10 사이 숫자 컴퓨터가 선정함.
		no1 := rand.Intn(10)
		no2 := rand.Intn(10)
		// 구구단 결과 값 (컴퓨터에 저장된 정답)
		res = no1 * no2
		fmt.Println("컴터가 문제를 냈습니다.")
		// 컴퓨터가 문제 내기
		fmt.Printf("%d x %d 는?\n", no1, no2)
		fmt.Println("답을 입력하세요!")
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
		user, _ := strconv.Atoi(line)
		// 유저가 적은 답이 컴퓨터에 저장된 정답과 같으면
		if user == res {
			fmt.Println("정답입니다!")
			// 시도 횟수 증가
			try++
			// 점수 증가
			grade++
			// 유저가 적은 답이 컴퓨터에 저장된 정답과 틀리면
		} else {
			fmt.Println("틀렸습니다.")
			// 시도 횟수 증가
			try++
			// 점수 감소
			grade--
		}
		// 점수 5점이 되었을 때
		if grade == 5 {
			fmt.Println("5점을 얻었습니다.")
			fmt.Println("게임이 끝납니다! 수고하셨습니다!")
			// 무한 반복문 탈출
			break
		}
	}
	// 결과 출력
	fmt.Printf("시도한 횟수 %d 만에 승리하셨습니다! \n", try)
}
