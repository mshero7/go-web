package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No Users")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	// Post Test (Create)
	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))

	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// Get Test
	users := new(Users)
	err = json.NewDecoder(res.Body).Decode(users)
	assert.NoError(err)
	assert.NotEqual(0, users.Id)

	id := users.Id
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users2 := new(Users)
	err = json.NewDecoder(res.Body).Decode(users2)
	assert.NoError(err)
	assert.Equal(users.Id, users.Id)
	assert.Equal(users.FirstName, users2.FirstName)
}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:1")

	// Post Test (Create)
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// Get Test
	users := new(Users)
	err = json.NewDecoder(res.Body).Decode(users)
	assert.NoError(err)
	assert.NotEqual(0, users.Id)

	id := users.Id
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), fmt.Sprintf("Delete User Id:%d", id))
}

// 값이 있을땐 바꿔줘야하고, 없을땐 어떻게 할것인가! 여기선 1번
// 1. Update Or Create
// 2. Update > Error
func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	userId := 1
	req, _ := http.NewRequest("PUT", ts.URL+"/users/"+strconv.Itoa(userId), strings.NewReader(`{"id":1, "first_name":"sangsu_updated", "last_name":"moon_updated", "email":"mshero7_updated@naver.com"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), fmt.Sprintf("No User Id:%d", userId))

	// Post Test (Create)
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// Get Test
	users := new(Users)
	err = json.NewDecoder(res.Body).Decode(users)
	assert.NoError(err)
	assert.NotEqual(0, users.Id)

	updateStr := fmt.Sprintf(`{"id":%d, "first_name":"sangsu_updated", "last_name":"moon_updated", "email":"mshero7_updated@naver.com"}`, userId)
	req, _ = http.NewRequest("PUT", ts.URL+"/users/"+strconv.Itoa(userId), strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updateUser := new(Users)
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(users.Id, updateUser.Id)
	assert.Equal("sangsu_updated", updateUser.FirstName)
	assert.Equal("moon_updated", updateUser.LastName)
	assert.Equal("mshero7_updated@naver.com", updateUser.Email)
}

func TestUsers_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewRestApiHandler())
	defer ts.Close()

	// Post Test (Create)
	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// Post Test (Create)
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	// Post Test (Create)
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangsu", "last_name":"moon", "email":"mshero7@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	res, err = http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	users := []*Users{}
	err = json.NewDecoder(res.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(3, len(users))
}
