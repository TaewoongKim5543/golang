package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // handlefunc는 요청에 따른 핸들러 등록.
		fmt.Fprint(w, "Hello World") // W값에 hello world를 주어 프린팅해라
	})

	http.ListenAndServe(":3000", nil) //요 함수를 통해 웹서버 간단제작. / nil은 웹구동을 시작하라는 의미 빠른 부팅을 위해.
}
