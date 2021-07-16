// package main

// import "fmt"

// func main() {

// 	var a int64
// 	var b int64

// 	a = 3
// 	b = 4

// 	fmt.Println(a, b)

// 	c := 4
// 	d := 5

// 	fmt.Println(c, d)
// }
package main // interface 방식으로 변경한 햄버거 만들기~

import (
	"fmt"
)

type MakeMenu interface {
	String() string
}
type SetMenu interface {
	case1() MakeMenu
}
type fuck struct {
	val string
}

func (f *fuck) MakingSetMenu(setmenu SetMenu) {
	aset := setmenu.case1()
	f.val += aset.String()
}

func (f *fuck) Strting() string {
	return "[Original buger" + f.val
}

func main() {
	set := &fuck{}
	fmt.Println(set)
}

///////////////////////////////////////////////////
