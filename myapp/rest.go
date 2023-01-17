package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") // Writer에 print 하라
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func usersInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Id를 파싱해줌
	vars := mux.Vars(r)

	fmt.Fprint(w, "User Id:", vars["id"])
}

func NewRestApiHandler() http.Handler {
	mux := mux.NewRouter() // http (x) gorilla mux

	mux.HandleFunc("/", indexxHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", usersInfoHandler)
	return mux
}
