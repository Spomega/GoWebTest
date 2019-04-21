package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

//Users should be able to view list of Users
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("HTTP Status expected 200, got: %d", w.Code)
	}
}

//User Story - Users should be able to create a User entity
func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/users", createUser).Methods("POST")

	userJSON := `{"firstname":"shiju","lastname":"Varghese", "email":"shiju@xyz.com"}`

	req, err := http.NewRequest("POST", "/users", strings.NewReader(userJSON))

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 201 {
		t.Errorf("HTTP Status expected: 201,got:%d", w.Code)
	}
}

//User Story - The Email Id of a User entity should be unique
func TestUniqueEmail(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")

	userJSON := `{"firstname":"shiju","lastname":"Varghese", "email":"shiju@xyz.com"}`

	req, err := http.NewRequest("POST", "/users", strings.NewReader(userJSON))

	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Errorf("Bad request expected, got:%d", w.Code)
	}
}

func TestGetUsersClient(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")

	server := httptest.NewServer(r)
	defer server.Close()

	usersUrl := fmt.Sprintf("%s/users", server.URL)
	request, err := http.NewRequest("GET", usersUrl, nil)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 200")
	}
}

func TestCreateUserClient(t *testing.T) {
	r := mux.NewRouter()

	r.HandleFunc("/users", createUser).Methods("POST")
	server := httptest.NewServer(r)
	defer server.Close()

	usersUrl := fmt.Sprintf("%s/users", server.URL)
	fmt.Println(usersUrl)

	userJSON := `{"firstname":"Rosmi","lastname":"Shiju","email":"rose@xyz.com"}`

	request, err := http.NewRequest("POST", usersUrl, strings.NewReader(userJSON))

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("HTTP Status expected: 201, got: %d", res.StatusCode)
	}
}
