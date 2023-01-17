package main

import (
	"go-web/myapp"
	"net/http"
)

func main() {
	// public 폴더의 하위 파일들에 접근가능하게 해줌
	// http.Handle("/", http.FileServer(http.Dir("public")))

	// 웹서버 실행, request 를 기다리는 상태가 됌
	go http.ListenAndServe(":3001", myapp.NewRestApiHandler())
	// http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
