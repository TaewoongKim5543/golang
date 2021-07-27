package main

import (
	"WEB3/myapp" // <===별도의 패키지화 없이 패키지 생성가능 !!!!!!!!!
	"net/http"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler()) // mux대신 myapp.NewHttpHandler()
}
