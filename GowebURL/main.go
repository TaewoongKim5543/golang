package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` // 물결표시, 백쿼트를 입력해서 json형식으로 설명을 붙이는것/ go와 json을 형식을 맞추어 주는 것.
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, err)
		return
	}

	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)

	w.Header().Add("content-type", "application/json")
	w.Header().Add("testing", "DongDaegu Bigbuger")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
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
