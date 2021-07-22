package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // handlefunc는 요청에 따른 핸들러 등록.
		fmt.Fprint(w, "Hello World") // W값에 hello world를 주어 프린팅해라
	})
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) { // handlefunc는 요청에 따른 핸들러 등록.
		fmt.Fprint(w, "Hello bar") // bar 추가!
	})

	http.Handle("/foo", &fooHandler{}) // 인터페이스 등록 구현! foo 추가!  boo와 foo는 각각 경로로 적용 인퍼테이스로도 구현되고 위 처럼도 구현이 된다~

	http.ListenAndServe(":3000", nil) //요 함수를 통해 웹서버 간단제작. / nil은 웹구동을 시작하라는 의미 빠른 부팅을 위해.
}
