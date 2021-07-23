package myapp

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//--index Handler res.body로 render된 data를 비교검증부분--//
func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

//-------------2단계 테스팅--------------------//
func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

//----------3단계 테스팅------------------------//
func TestBarPathHandler_WitName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar?name=junyoung", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	//app.go페이지에서 index Handler에서 render data를 읽어오는 부분----//
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello junyoung!", string(data))
}

//------4단계 테스팅---------------------------//
func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/foo", nil)
	//mux로 "/"경로를 타켓 라우팅(분배) 응신 렌더링 시키는 부분
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code) //성공
}
