package myapp

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
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

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	// 실제 http 프로토콜을 사용하지 않는 http 테스트 패키지
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	// 실제 http 프로토콜을 사용하지 않는 http 테스트 패키지
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`),
	)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("sangsu", user.FirstName)
	assert.Equal("moon", user.LastName)
	assert.Equal("mshero7@naver.com", user.Email)
}

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)

	// 파일 읽기
	path := "C:/Users/ssmoon/Desktop/FirstView.jpg"
	file, err := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	// multiform 파일을 만듬
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/file/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "C:/Users/ssmoon/Desktop/Moon/go-web/uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath) // get file info
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)

	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
