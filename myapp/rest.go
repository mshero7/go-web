package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Users struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// 실무에선 UpdatedFirstName bool 과 같이 업데이트될 컬럼의 정보에 대한 T/F를 받음
type UpdateUsers struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var usersMap map[int]*Users
var lastId int

func indexxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") // Writer에 print 하라
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(usersMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}

	users := []*Users{}
	for _, u := range usersMap {
		users = append(users, u)
	}

	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func getUsersInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Id를 파싱해줌
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// map에서 key 를 통해 찾으면 값과 error 를 반환해줌
	user, ok := usersMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "appliation/json")

	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	users := new(Users)
	err := json.NewDecoder(r.Body).Decode(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastId++
	users.Id = lastId
	users.CreatedAt = time.Now()
	usersMap[users.Id] = users

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "appliation/json")

	data, _ := json.Marshal(users)
	fmt.Fprint(w, string(data))
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	_, ok := usersMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	// key에 해당하는 value를 map에서 삭제해줌
	delete(usersMap, id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete User Id:", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	updateUser := new(Users)
	err := json.NewDecoder(r.Body).Decode(updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	users, ok := usersMap[updateUser.Id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", updateUser.Id)
		return
	}

	if updateUser.FirstName != "" {
		users.FirstName = updateUser.FirstName
	}

	if updateUser.LastName != "" {
		users.LastName = updateUser.LastName
	}

	if updateUser.Email != "" {
		users.Email = updateUser.Email
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(users)
	fmt.Fprint(w, string(data))
}

func NewRestApiHandler() http.Handler {
	usersMap = make(map[int]*Users)
	mux := mux.NewRouter() // http (x) gorilla mux

	mux.HandleFunc("/", indexxHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUsersInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserHandler).Methods("DELETE")
	mux.HandleFunc("/users/{id:[0-9]+}", updateUserHandler).Methods("PUT")

	return mux
}
