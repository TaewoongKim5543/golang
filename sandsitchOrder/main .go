package main

import "fmt"

//-----------------------구조체------------------------------//
type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberry struct {
}

type Sandwitch struct {
	val string
}

//-----------------------오렌지잼 추가 구조체------------------------------//
type OrangeJam struct {
	opened bool
}
type SpoonOfOrangeJam struct {
}

//-----------------------함 수------------------------------//
func GetBreads(num int) []*Bread {
	breads := make([]*Bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}
func GetOneSpoon(_ *StrawberryJam) *SpoonOfStrawberry {
	return &SpoonOfStrawberry{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

//-----------------------오렌지잼 추가 함수------------------------------//
func OpenOrangeJam(jam *OrangeJam) {
	jam.opened = true
}
func GetOneOrangeJamSpoon(_ *OrangeJam) *SpoonOfOrangeJam {
	return &SpoonOfOrangeJam{}
}
func PutOrangeJamOnBread(bread *Bread, jam *SpoonOfOrangeJam) {
	bread.val += " + Orange Jam"
}

//-----------------------오렌지잼 추가 함수------------------------------//

//-----------------------실행순서------------------------------//
func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)

	// jam := &StrawberryJam{}

	// 2. 딸기잼 뚜껑을 연다.
	// OpenStrawberryJam(jam)

	//-----------------------오렌지잼 추가 순서 ------------------------------//
	jam := &OrangeJam{}
	// 2-1. 오렌지잼 뚜겅을 연다.
	OpenOrangeJam(jam)
	spoon := GetOneOrangeJamSpoon(jam)
	PutOrangeJamOnBread(breads[0], spoon)
	//-----------------------오렌지잼 추가 순서 ------------------------------//

	// 3. 딸기잼을 한스푼 뜬다.
	// spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	// PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
