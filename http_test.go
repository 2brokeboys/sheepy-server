package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/2brokeboys/sheepy-server/db"
	"github.com/stretchr/testify/assert"
)

func postTest(t *testing.T, path string, body string, code int, resp string) {
	r := setupRouter()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", path, strings.NewReader(body))

	r.ServeHTTP(res, req)
	assert.Equal(t, code, res.Code)
	assert.JSONEq(t, resp, res.Body.String())
}

func TestNoLogin(t *testing.T) {
	for _, path := range []string{"/newGame", "/queryUser"} {
		postTest(t, path, "", 403, `{"error":"not logged in"}`)
	}
}

func TestLogin(t *testing.T) {
	db.InitDB()
	postTest(t, "/login", "", 400, `{"error":"invalid data"}`)
	postTest(t, "/login", `{"user":"a","password":"a"}`, 401, `{"error":"invalid credentials"}`)
	postTest(t, "/login", `{"user":"test","password":"a"}`, 200, `{"success":true}`)
}
