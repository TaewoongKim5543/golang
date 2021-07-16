package main

import "fmt"

//-------------------------------Interface 정의-------------------------------//
type MakeSashimi interface {
	String() string
}

type SashimiSpecies interface {
	AddMakedSashimi() MakeSashimi
}

//-------------------------------모듬회 오브젝트1 정의-------------------------------//
type AssortedSashimi struct {
	sashimi string
}

func (a *AssortedSashimi) Addsashimispecies(sashimispecies SashimiSpecies) { // filling(속)은 종류가 무관하다.
	onesashimispecies := sashimispecies.AddMakedSashimi()
	a.sashimi += onesashimispecies.String()
}
func (a *AssortedSashimi) String() string { //Bread 매서드2 string
	return "[ Busan Haeundae AssortedSashimi" + a.sashimi + " ]"
}

//-------------------------------횟감1(도다리) 오브젝트2 정의-------------------------------//
type Flounder struct {
}

func (f *Flounder) AddMakedSashimi() MakeSashimi {
	return &MakeFlounderSashimi{}
}

//-------------------------------회만들기1(도다리회) 오브젝트2 정의-------------------------------//
type MakeFlounderSashimi struct{}

func (f *MakeFlounderSashimi) String() string {
	return " = flounder Sashimi"
}

//-------------------------------횟감2(광어) 오브젝트2 정의-------------------------------//
type Flatfish struct {
}

func (f *Flatfish) AddMakedSashimi() MakeSashimi {
	return &MakeFlatfishSashimi{}
}

//-------------------------------회만들기2(광어회) 오브젝트2 정의-------------------------------//
type MakeFlatfishSashimi struct{}

func (f *MakeFlatfishSashimi) String() string {
	return " + Flatfish Sashimi"
}

//-------------------------------실행 부분-------------------------------//
func main() {
	assortedSashimi := &AssortedSashimi{}
	sashimispecies1 := &Flounder{}
	sashimispecies2 := &Flatfish{}

	assortedSashimi.Addsashimispecies(sashimispecies1)
	assortedSashimi.Addsashimispecies(sashimispecies2)

	fmt.Println(assortedSashimi)
}
