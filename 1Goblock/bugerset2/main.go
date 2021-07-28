package main

import "fmt"

type TongsOfPatty interface { //외부 공개 메서드 = 관계
	String() string //PutOfPatty은 String()이라는 외부 공개 메서드
}

type Patty interface { //오로지 관계만 선언해본다
	GetOneTongs() TongsOfPatty //Jam은 GetOnsSpoon()이라는 외부 공개 메서드
}

type BigBurger struct { //Bread 객체를 만든다
	val string
}

func (b *BigBurger) PutPatty(patty Patty) { //오렌지잼이든 스토우베리 잼이든 상관없다.
	tongs := patty.GetOneTongs()
	b.val += tongs.String()
}

func (b *BigBurger) String() string { //Bread 메서드2 String
	return "[대구 모듬회 = " + b.val + "]"
}

type ChickenPatty struct {
}

func (j *ChickenPatty) GetOneTongs() TongsOfPatty {
	return &TongsOfChickenPatty{}
}

type SlicePatty struct {
}

func (j *SlicePatty) GetOneTongs() TongsOfPatty {
	return &TongsOfSlicePatty{}
}

type CheesePatty struct {
}

func (j *CheesePatty) GetOneTongs() TongsOfPatty {
	return &TongsOfCheesePatty{}
}

type TongsOfChickenPatty struct {
}

func (s *TongsOfChickenPatty) String() string {
	return " +  도다리회"
}

type TongsOfSlicePatty struct {
}

func (s *TongsOfSlicePatty) String() string {
	return " + 광어회"
}

type TongsOfCheesePatty struct {
}

func (s *TongsOfCheesePatty) String() string {
	return "오징어회"
}

func main() {
	BigBurger := &BigBurger{}
	//jam := &StrawberryJam{}
	//jam := &OrangeJam{}
	patty := &CheesePatty{}
	patty2 := &ChickenPatty{}
	patty3 := &SlicePatty{}

	BigBurger.PutPatty(patty)
	BigBurger.PutPatty(patty2)
	BigBurger.PutPatty(patty3)

	fmt.Println(BigBurger)
}
