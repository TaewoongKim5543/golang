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

	// user에 담겨있는 json 데이터들을 json.Marshal로 병합(Marshal: 한국어로 번역하면 '합치다' 정도)
	// json.Marshal : 매개변수 - 인터페이스, 리턴값 - 바이트 배열
	data, _ := json.Marshal(user)

	// w.Header().Add : 리스폰스 헤더 내용 작성
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(http.StatusOK)

	// json.Marshal 리턴값으로 저장된 변수 data를 스트링화 하고, http.ResponseWriter로 작성.
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// Fprintf - 바디(본문) 내용 렌더링
	// 입력되는 값이 없는 경우 디폴트로 Hellow World! 출력
	fmt.Fprintf(w, "Hello %s!", name)

	// 응용 케이스 : fmt.Fprintf 로 본문 내용 변경 후 app_test.go에서 바꾼 내용으로 문자열 비교 테스트해보기.
	// fmt.Fprintf(w, "나를 똑같이 맞춰줘!")
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux() //라우터 클래스 만든다. mux 등록
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler) //mux

	mux.Handle("/foo", &fooHandler{}) // mux
	return mux
}
