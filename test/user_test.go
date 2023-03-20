package test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
)

func clearUserTable(test Test) {
	test.clearTable("TRUNCATE TABLE users")
}

func registerUser(i int) {
	gettest := Test{}
	gettest.initTest()

	for j := 0; j < i; j++ {
		user_id := uuid.New()
		user_name := "user " + strconv.Itoa(j)
		user_email := "user" + strconv.Itoa(j) + "@gmail.com"
		user_role := "user"
		created_at := time.Now().Format("2006-01-02 15:04:05")

		gettest.App.DB.Exec("INSERT INTO users VALUES(?, ?, ?, ?, ?, NULL);", user_id, user_name, user_email, user_role, created_at)
	}
}

func TestRegisterUser(t *testing.T) {
	gettest := Test{}
	gettest.initTest()

	clearUserTable(gettest)
	var jsonStr = []byte(`{
			"username" : "Ramadhan",
			"email" : "ramadhan@gmail.com",
			"userRole" : "user"
		}`)
	req, _ := http.NewRequest("POST", "/v1/register", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := gettest.executeRequest(req)

	gettest.checkResponseCode(t, http.StatusOK, response.Code)
}

func TestLoginUser(t *testing.T) {
	gettest := Test{}
	gettest.initTest()

	clearUserTable(gettest)
	registerUser(1)
	var jsonStr = []byte(`{
			"email" : "user1@gmail.com",
			"password" : "password1"
		}`)
	req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	response := gettest.executeRequest(req)

	gettest.checkResponseCode(t, http.StatusOK, response.Code)
}
