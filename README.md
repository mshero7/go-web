# go-web

부족했던 웹 개념도 채우고, go로 웹도 개발해보고 go 테스트에 익숙해지기 위해 진행

주석과 app.go에 handler 를 우겨넣는 삽질중

https://www.youtube.com/@TuckerProgramming 의 Go 로 만드는 웹 참고.

Note
1. New(Type)은 포인터 Type을 반환함
2. Map[key] 할때 value, err 을 리턴함
3. Go에서도 SPA같이 html 파일을 템플릿처럼 사용해줄수있따..! 내장라이브러리인 template를 통해 비슷하게 가능하고, 외부라이브러리인 render(렌더), pat(좀더 쉽고 이쁘게 사용 가능한 gorilla Router)를 활용할 수 있다. negori 는 미들웨어.
