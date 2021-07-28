package main // 수 예제 go mastering 책

import (
	"fmt"
)

func main() {
	c1 := 12 + 1i                      // 복소수 a + bi 형태로 표현해줌 , a, b는 실수, i는 x^2 = -1 의 해
	c2 := complex(5, 7)                //complex64 => 실수, 허수부 표현에 float32 2개 사용/ complex128 => 실수, 허수부 표현에 float64 2개 사용
	fmt.Printf("Type of c1: %T\n", c1) // ------1. 출력 c1의 타입
	fmt.Printf("Type of c1: %T\n", c2) // ------2. 출력 c2의 타입

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)             // ------3. 출력 c3 값
	fmt.Printf("Type of c3: %T\n", c3) // ------4. 출력 c3의 타입

	cZero := c3 - c3
	fmt.Println("cZero:", cZero) // ------5. 출력 cZero 값

	// 복소수를 이용해서 계산해주는 메인함수, c1,c2는 복소수를 직접 만듬, c3는 cZero처럼 기존 복소수로 계산해 간접적으로 만듬
	// ** acomplex := 12 + 2* i 안하는 이유 -> 2개의 의도하지 않은 결과 나옴 / Go 입장에서 이문장은 덧셈과 곱셈을 하는 것임.
	// 1-현재 스코프에 숫자를 담은 변수 i가 없으면 구문에러 발생, 컴파일 실패함. 2-변수 i가 있다면 컴파일 에러는 없어도 버그발생.

	x := 12
	k := 5
	fmt.Println(x)                   // ------6. 출력 X 값
	fmt.Printf("Type of x: %T\n", x) // ------7. 출력 X 의 타입

	div := x / k
	fmt.Println("div", div) // ------8. 출력 x / k (정수나누기결과)
	// 부호 있는 정수 사용, 여기서 Go는 두 정수의 나눗셈을 보고 결과를 정수로 보내주므로 나머지가 날라가버린다.
	// 부동소수점 수를 정수로 변환 시 소수점 자리 숫자 날림. == 데이타 손실 발생

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n) // ------9. 출력 m, n 값

	y := 4 / 2.3
	fmt.Println("y:", y)             // ------10. 출력 y 값
	fmt.Printf("Type of y: %T\n", y) // ------11. 출력 y 의 타입 임마는 왜 플로트고...

	divFloat := float64(x) / float64(k) // float64(x) / k -->이러면 타입 맞추라고 에러 출력.
	fmt.Println("divFloat", divFloat)
	fmt.Printf("Type of divFloat: %T\n", divFloat)
	// 부동소수점 다루는법, float64로 두 정수를 나눌때 부동소수점 수를 생성하게 만듬.
}
