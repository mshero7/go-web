package myapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	// 실제 http 프로토콜을 사용하지 않는 http 테스트 패키지
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)

	assert.Equal("Hello World", string(data))
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	// 실제 http 프로토콜을 사용하지 않는 http 테스트 패키지
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	// req 를 인덱스(/) 로 걸면 함수를 직접호출하기때문에 테스트가 패스되는 오류가 있음
	// barHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)

	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	// 실제 http 프로토콜을 사용하지 않는 http 테스트 패키지
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=sangsu", nil)

	// req 를 인덱스(/) 로 걸면 함수를 직접호출하기때문에 테스트가 패스되는 오류가 있음
	// barHandler(res, req)

	// mux 통해서 등록된 핸들러 정보들로 테스트하게 처리.
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)

	assert.Equal("Hello sangsu!", string(data))
}
