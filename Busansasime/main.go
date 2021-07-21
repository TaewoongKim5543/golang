package main

import "fmt"

//-------------------------------Interface 정의-------------------------------//
type Dishofsashomi interface { // 외부 공개 메서드 1 문자열을 받아 동작하는 스트링함수
	String() string // 스트링함수기능을 Dishofsashomi 인터페이스 지원
}

type Slice interface { // 외부 공개 메서드 1 디시오브 사시미를 받아서 동작하는
	Getonedish() Dishofsashomi // 겟원 디시함수기능을 slice 인터페이스로 지원
}

//-------------------------------모듬회 오브젝트1 정의-------------------------------//
type Sashimiset struct { //모듬회 구조체는 문자열 sasimi 라는 기본값.
	sashimi string
}

//모듬회 구조체 함수 애드 슬라이스기능은 슬라이스 라는 인터페이스 값을 받아 동작
func (a *Sashimiset) Addslice(slice Slice) {
	oneslice := slice.Getonedish() //oneslice
	a.sashimi += oneslice.String()
}
func (a *Sashimiset) String() string {
	return "[ Daegu Sashimiset" + a.sashimi + " ]"
}

//-------------------------------횟감1(도다리) 오브젝트2 정의-------------------------------//
type Flounder struct {
}

func (f *Flounder) Getonedish() Dishofsashomi {
	return &MakeFlounderSashimi{}
}

//-------------------------------회만들기1(도다리회) 오브젝트2 정의-------------------------------//
type MakeFlounderSashimi struct{}

func (f *MakeFlounderSashimi) String() string {
	return " = flounder Sashimi"
}

//-------------------------------횟감2(광어) 오브젝트3 정의-------------------------------//
type Flatfish struct {
}

func (f *Flatfish) Getonedish() Dishofsashomi {
	return &MakeFlatfishSashimi{}
}

//-------------------------------회만들기2(광어회) 오브젝트3 정의-------------------------------//
type MakeFlatfishSashimi struct{}

func (f *MakeFlatfishSashimi) String() string {
	return " + Flatfish Sashimi"
}

//-------------------------------실행 부분-------------------------------//
func main() {
	sashimiset := &Sashimiset{}
	slice1 := &Flounder{}
	slice2 := &Flatfish{}

	sashimiset.Addslice(slice1)
	sashimiset.Addslice(slice2)

	fmt.Println(sashimiset)
}
