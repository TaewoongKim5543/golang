package myapp

import (
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

	indexHandler(res, req)

	assert.Equal(http.StatusOK, res.Code) //( , ) 앞 뒤로 같냐 다르냐를 확인 점검 후 뿌려줌 assert
	data, _ := ioutil.ReadAll(res.Body)   //데이터를 스트링으로 읽어옴. 응답된 바디 본문에서!,
	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder() //실제 response를 사용하지 않고 테스팅 응신방법
	req := httptest.NewRequest("GET", "/bar", nil)

	barHandler(res, req)

	assert.Equal(http.StatusOK, res.Code) //( , ) 앞 뒤로 같냐 다르냐를 확인 점검 후 뿌려줌 assert
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data)) // 얘가 go 피일의 본문과 동일하게 받아서 확인 후 틀리면 X
}
