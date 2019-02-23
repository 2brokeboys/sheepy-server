package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/2brokeboys/sheepy-server/common"
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
	db.AddTestUsers(t)
	postTest(t, "/login", "", 400, `{"error":"invalid data"}`)
	postTest(t, "/login", `{"username":"a","password":"a"}`, 401, `{"error":"invalid credentials"}`)
	postTest(t, "/login", `{"username":"foo","password":"12456"}`, 401, `{"error":"invalid credentials"}`)
	postTest(t, "/login", `{"username":"foo","password":"123456"}`, 200, `{"success":true}`)
}

func TestSession(t *testing.T) {
	db.InitDB()
	db.AddTestUsers(t)

	// Setup test HTTP server
	ts := httptest.NewServer(setupRouter())
	defer ts.Close()

	// Setup HTTP client
	jar, err := cookiejar.New(nil)
	assert.Nil(t, err)
	c := &http.Client{
		Jar: jar,
	}

	pt := func(path string, body string, code int, resp string) {
		res, err := c.Post(ts.URL+path, "", strings.NewReader(body))
		assert.Nil(t, err)
		assert.Equal(t, code, res.StatusCode)
		b, _ := ioutil.ReadAll(res.Body)
		assert.JSONEq(t, resp, string(b))
	}

	// Login with valid credentials
	pt("/login", `{"username":"foo","password":"123456"}`, 200, `{"success":true}`)

	// Second login should fail
	pt("/login", `{"username":"foo","password":"123456"}`, 409, `{"error":"already logged in"}`)

	// Create new game with invalid data
	pt("/newGame", "", 400, `{"error":"invalid data"}`)

	// Create new game
	g := &common.Game{
		Participants: [4]int{0, 1, 2, 3},
		Player:       1,
		Playmate:     -1,
		GameType:     common.SoloEichel,
		Points:       120,
		Schwarz:      true,
		Runners:      5,
		Virgins:      0,
	}
	b, err := json.Marshal(g)
	assert.Nil(t, err)
	pt("/newGame", string(b), 200, `{"success":true}`)

	// Query users - ok
	pt("/queryUser", `{"search":"oo"}`, 200, `{"success":true,
	"users":[{"id":1, "username":"foo", "name":""},{"id":3, "username":"moo", "name":""}]}`)

	// Query users - fail
	pt("/queryUser", "", 400, `{"error":"invalid data"}`)

	// Query users - empty
	pt("/queryUser", `{"search":"sfsefasdf"}`, 200, `{"success":true,"users":[]}`)
}
