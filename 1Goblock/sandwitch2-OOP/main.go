package main // OOP방식으로 변경한 sandwitch 만들기~

import "fmt"

type Bread struct { // 빵 객체 생성.
	val string
}

type StrawberryJam struct {
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string { // SpoonOfStrawberryJam의 1개 매소드 string
	return "+ strawberryjam"
}

func (j *StrawberryJam) GetOneSpoon() *SpoonOfStrawberryJam { // SpoonOfStrawberryJam의 1개 매소드 GetOneSpoon
	return &SpoonOfStrawberryJam{}
}

func (b *Bread) PutJam(jam *StrawberryJam) { //Bread 매서드1 PutJam
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string { //Bread 메서드2 String
	return "bread " + b.val + " + bread"
}

func main() {
	bread := &Bread{}
	jam := &StrawberryJam{}
	bread.PutJam(jam)

	fmt.Println(bread)
}
