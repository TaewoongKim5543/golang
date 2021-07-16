package main

import "fmt"

func main() {
	var a int   // 컴퓨터 지원하는 int64 범위
	var b int8  // 2^8/2 -128~127
	var c int16 // 2^16/2 -32768~32767
	var d int32
	var e int64

	a = 1000000000000000000
	b = -127
	c = 32767
	d = 2000000000
	e = 9000000000000000000

	fmt.Println(a, b, c, d, e)

	//unsugned
	var f uint
	var g uint8
	var h uint16
	var i uint32
	var j uint64

	f = 999999999999999999
	g = 255
	h = 65353
	i = 4000000000
	j = 9000000000000000000

	fmt.Println(f, g, h, i, j)

	var s string
	var t string
	var r rune
	var o bool

	s = "Go Block"
	t = "Go"
	r = 'G'
	o = true

	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(r)
	fmt.Println(string(r))
	fmt.Println(o)
}
