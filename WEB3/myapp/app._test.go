package myapp

import (
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
	assert.Equal(http.StatusOK, res.Code)
}
