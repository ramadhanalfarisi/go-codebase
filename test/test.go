package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ramadhanalfarisi/go-codebase-kocak/app"
)

type Test struct{
	App app.App
}

func (test *Test) initTest() {
	test.App.ConnectDB()
	test.App.Routes()
}

func (test *Test) clearTable(query string) {
	test.App.DB.Exec(query)
}

func (test *Test) checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func (test *Test) executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	test.App.Router.ServeHTTP(rr, req)

	return rr
}