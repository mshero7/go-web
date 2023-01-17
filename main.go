package main

import (
	"go-web/myapp"
	"net/http"
)

func main() {
	// 웹서버 실행, request 를 기다리는 상태가 됌
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
