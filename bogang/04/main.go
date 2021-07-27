package main //시험 점수 메기는 함수 만들기
import "fmt"

func main() {
	// 학생 시험 점수 메기는 함수 만들어 보자

	// 1. ex 과학 80점, 수학 60점, 국어 60점 일때 평균 점수 출력함수
	e, f, g := examSubject("과학", "수학", "영어")
	a, b, c := examGrade(70, 80, 70)

	avg := avg(a, b, c)
	printAVG(avg, e, f, g)
}

// 과목점수와 과목을 3개를 받는 함수를 만들어보자
func examSubject(r1, r2, r3 string) (string, string, string) {
	return r1, r2, r3
}

func examGrade(o1, o2, o3 int) (int, int, int) {
	return o1, o2, o3
}

// 과목과 과목 점수를 가져와 평균을 구하는 함수를 만들어 보자
func avg(o1, o2, o3 int) float64 {
	var avg float64
	var sum int
	sum = o1 + o2 + o3

	avg = float64(sum / 3)

	return avg

}

func printAVG(avg float64, a, b, c string) {
	fmt.Printf("평균점수는 %f", avg)
	fmt.Printf("\n시험 본 과목은  %s %s %s이다.", a, b, c)

}
