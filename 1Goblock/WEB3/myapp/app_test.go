package myapp // app.go 의 테스트 파일 !

import (
	"fmt"
	"io/ioutil" //에러값과 리턴값 2개를 받아주는 역할
	"net/http"
	"net/http/httptest"
	"testing" // 얘는 테스트 파일이에용

	"github.com/stretchr/testify/assert"
)

//--index Handler res.body로 render된 data를 비교검증부분--//
func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 네트워크/ response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/", nil)

	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // <============GO interface 기능 적용

	// if res.Code != http.StatusOK {
	// 	t.Fatal("Failed!!", res.Code)
	// }

	fmt.Println("1단계 테스팅 성공")

	//stretchr의 testify에서 가져온 assert패키지

	assert.Equal(http.StatusOK, res.Code) //equal 안의 두개 인자 유사성 비교
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data)) //equal 안의 두개 인자 유사성 비교  헬로 월드와 string data
}

//-------------2단계 테스팅--------------------//
func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	fmt.Println("2단계 테스팅 성공")

	assert.Equal(http.StatusOK, res.Code)
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

//----------3단계 테스팅------------------------//
func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar?name=TaeWoong", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	fmt.Println("3단계 테스팅 성공")

	assert.Equal(http.StatusOK, res.Code)
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello TaeWoong!", string(data))
}

//------4단계 테스팅---------------------------//
func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/foo", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	fmt.Println("4단계 테스팅 성공")

	assert.Equal(http.StatusBadRequest, res.Code) //성공
}

//------5단계 테스팅---------------------------//
func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/foo", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	fmt.Println("5단계 테스팅 성공")

	assert.Equal(http.StatusBadRequest, res.Code) //성공
}
