package myapp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/", nil)

	// indexHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	fmt.Println("왜안되노..")
	assert.Equal(http.StatusOK, res.Code) // 잘 돌아가는지 아닌지 검사 .

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=TaeWoong", nil)

	// barHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	// data, _ := ioutil.ReadAll(res.Body) : 본문 내용 가지고 와서 data 변수에 저장
	// ioutil.ReadAll - 리턴값 2개 : []byte 와 error 값
	// res.Body : response body 변수 네이밍
	data, _ := ioutil.ReadAll(res.Body)
	// data에 저장된 바이트 배열을 스트링으로 바꾸고, 기준이 되는 문자열 "Hello World!"와 비교
	assert.Equal("Hello TaeWoong!", string(data))
}
