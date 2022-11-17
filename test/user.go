package test

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ramadhanalfarisi/go-codebase-kocak/app"
)

var a app.App

func TestMain(m *testing.M) {
	a.ConnectDB()

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func clearTable() {
	a.DB.Exec("DELETE FROM users;")
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func Hashing(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func registerUser(i int) {
	for j := 0; j < i; j++ {
		user_id := uuid.New()
		user_name := "user " + strconv.Itoa(j)
		user_email :="user" + strconv.Itoa(j) + "@gmail.com"
		user_role := "user"
		created_at := time.Now().Format("2006-01-02 15:04:05")
	
		a.DB.Exec("INSERT INTO users VALUES(?, ?, ?, ?, ?, NULL);",user_id,user_name,user_email,user_role,created_at)
	}
}

func TestRegisterUser(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{
			"username" : "Ramadhan",
			"email" : "ramadhan@gmail.com",
			"userRole" : "user"
		}`)
	req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestLoginUser(t *testing.T) {
	clearTable()
	registerUser(1)
	var jsonStr = []byte(`{
			"email" : "user1@gmail.com",
			"password" : "password1"
		}`)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
