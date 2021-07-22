package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Helo bar!")
}

func main() {
	mux := http.NewServeMux() //Mux 라우터를 통한 구현! Mux 클래스를 만들어주라
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", barHandler) //mux

	mux.Handle("/foo", &fooHandler{}) //mux

	http.ListenAndServe(":3000", mux) // mux 라우트 등록
	//요 함수를 통해 웹서버 간단제작.
}
