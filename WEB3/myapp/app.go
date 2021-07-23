package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` //어노테이션)(Annotaion) 설명을 붙히는것
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprint(w, "Hello World") // Fprint은 w값에 "Hello World" 값을 주어서 프린팅해라는 의미
}

type fooHandler struct{} //인스턴스 생성

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //ServeHTTP 인터페이스 구현
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) //인터페이스, JSON형식 결과값은 바이트 어레이
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name) // 본문, 바디에
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() //라우터 클래스 만든다. mux 등록
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler) //mux

	mux.Handle("/foo", &fooHandler{}) // mux
	return mux
}
