package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AashrayAnand/tripit/models"
)

// register the new user specified by the request parameters
func register(params string, t *testing.T, expectedCode int) {
	req, _ := http.NewRequest("POST", "/user/create", strings.NewReader(params))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		status := w.Code
		return status == expectedCode
	})
}

// login the user specified by the request parameters
func login(params string, t *testing.T, expectedCode int) {
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(params))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// serve request again and check that it fails with 400
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		status := w.Code
		decoder := json.NewDecoder(w.Body)
		var resp models.AuthResp
		_ = decoder.Decode(&resp)
		// if login was supposed to succeed, check if auth token received
		if expectedCode == http.StatusOK {
			return status == expectedCode && resp.Auth != ""
		}

		// if login was not expected to succceed, check we received
		// the appropriate status code
		return status == expectedCode
	})
}

// test user registration works properly
func TestRegisterNewUser(t *testing.T) {
	params := getUserParameters()
	// register and check it was successful
	register(params, t, http.StatusOK)
}

// test that duplicate registration is disallowed
func TestRegisterDuplicateUser(t *testing.T) {
	params := getUserParameters()
	// register and check it was successful
	register(params, t, http.StatusOK)
	// attempt to register again, check it failed
	register(params, t, http.StatusBadRequest)
}

// test user registration + login works properly
func TestRegisterAndLogin(t *testing.T) {
	params := getUserParameters()
	// register user, check it was successful
	register(params, t, http.StatusOK)
	// login user, check it was successful
	login(params, t, http.StatusOK)
}
