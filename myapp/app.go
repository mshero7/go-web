package myapp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type fooHandler struct{}

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json") // json format에 맞게 처리되게끔
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") // Writer에 print 하라
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// Request 에 담긴 정보들통해 get param Argument
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "World"
	}

	fmt.Fprintf(w, "Hello %s!", name)
}

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")

	r.Header.Add("Content-Type", "multipart/form-data")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	dirname := "C:/Users/ssmoon/Desktop/Moon/go-web/uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func NewHttpHandler() http.Handler {
	// static HTML FILE
	fileServerHtml := http.FileServer(http.Dir("./public"))

	// mux 라는 라우터 인스턴스를 통해 처리
	mux := http.NewServeMux()

	// 초기경로, 절대경로 도메인의 첫번째 경로
	mux.HandleFunc("/", indexHandler)

	// 함수를 외부로 빼서 넣는것도 가능.
	mux.HandleFunc("/bar", barHandler)

	// 핸들러를 인스턴스 형태로 등록할때는 Handle()함수
	mux.Handle("/foo", &fooHandler{})

	// 파일 업로드용 핸들러
	mux.Handle("/file/", http.StripPrefix("/file/", fileServerHtml))
	mux.HandleFunc("/file/uploads", uploadsHandler)

	return mux
}
