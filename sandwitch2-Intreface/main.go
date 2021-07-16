package main // interface 방식으로 변경한 sandwitch 만들기~

import "fmt"

/////////////////////////////////////////////////////////////////////
type SpoonofJam interface { //외부 메서드 = 관계
	String() string // SpoonOfJam은 String()이라는 외부공개메서드이다.
}

type Jam interface { //오로지 관계만 선언한다.
	GetSpoonofJam() SpoonofJam //Jam은 GetOneSpoon() 이라는 외부 공개 메서드이다.
}

type Bread struct { // bread 객체를 만든다
	val string
}

func (b *Bread) MakingSetMenu(jam Jam) { // Jam은 종류가 무관하다.
	onespoon := jam.GetSpoonofJam()
	b.val += onespoon.String()
}

func (b *Bread) String() string { //Bread 매서드2 string
	return "bread1" + b.val
}

type StrawberryJam struct{}

func (j *StrawberryJam) GetOneSpoon() SpoonofJam {
	return &SpoonOfStrawberryJam{}
}

// type OrangeJam struct{}

// func (j *OrangeJam) GetOneSpoon() SpoonofJam {
// 	return &SpoonOfOrangeJam{}
// }

// type AppleJam struct{}

// func (j *AppleJam) GetOneSpoon() SpoonofJam {
// 	return &SpoonOfAppleJam{}
// }

type SpoonOfStrawberryJam struct{}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ strawberry"
}

// type SpoonOfOrangeJam struct{}

// func (s *SpoonOfOrangeJam) String() string {
// 	return "+ Orange"
// }

// type SpoonOfAppleJam struct{}

// func (s *SpoonOfAppleJam) String() string {
// 	return "+ Apple"
// }

func main() {
	bread := &Bread{}

	// jam := &StrawberryJam{}
	// // //jam := &OrangeJam{}
	// // jam := &AppleJam{}
	// bread.MakingSetMenu(jam)

	fmt.Println(bread)
}
