package main

import "go-web/oauth"

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

// func main() {
// 	// public 폴더의 하위 파일들에 접근가능하게 해줌
// 	// http.Handle("/", http.FileServer(http.Dir("public")))

// 	// 웹서버 실행, request 를 기다리는 상태가 됌
// 	// go http.ListenAndServe(":3001", myapp.NewRestApiHandler())
// 	// go http.ListenAndServe(":3002", myapp.NewDecoServer())
// 	// http.ListenAndServe(":3000", myapp.NewHttpHandler())

// 	// Go 내장 라이브러리(template) struct 를 특정 output으로 보내줄수있음
// 	user := User{Name: "tucker", Email: "tucker@naver.com", Age: 25}
// 	// 1. 파싱 string으로 templdate 사용
// 	tmplate, err := template.New("Template1").Parse("Name: {{.Name}}\nEmail: {{.Email}}\nAge: {{.Age}}")
// 	// 2. template file을 사용
// 	tmplate2, err := template.New("Template1").ParseFiles("templates/template1.tmpl")

// 	if err != nil {
// 		panic(err)
// 	}

// 	tmplate.Execute(os.Stdout, user)
// 	tmplate2.ExecuteTemplate(os.Stdout, "template1.tmpl", user)
// }

func main() {
	oauth.OauthServerExec()
}
