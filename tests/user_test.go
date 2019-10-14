package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/util"
	"github.com/gin-gonic/gin"
)

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}

// get the payload for login/create requests
func getUserParameters() string {
	parameters := url.Values{}
	// generate unique profile info, to avoid
	// issues with collision with database
	// without having to clear db
	parameters.Add("user", session.RandomString(10))
	parameters.Add("name", session.RandomString(10))
	parameters.Add("pass", session.RandomString(10))
	return parameters.Encode()
}

func register(params string, t *testing.T, expectedCode int) {
	req, _ := http.NewRequest("POST", "/user/create", strings.NewReader(params))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		status := w.Code
		return status == expectedCode
	})
}
func login(params string, t *testing.T, expectedCode int) {
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(params))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// serve request again and check that it fails with 400
	testHTTPResponse(t, router, req, func(w *httptest.ResponseRecorder) bool {
		status := w.Code
		return status == expectedCode
	})
}

// server to use for tests
var router = util.RunServer()

func TestRegisterNewUser(t *testing.T) {
	params := getUserParameters()
	register(params, t, http.StatusOK)
}

func TestRegisterDuplicateUser(t *testing.T) {
	params := getUserParameters()
	register(params, t, http.StatusOK)
	register(params, t, http.StatusBadRequest)
}

func TestRegisterAndLogin(t *testing.T) {
	params := getUserParameters()
	// REGISTER USER
	register(params, t, http.StatusOK)
	// LOGIN USER
	login(params, t, http.StatusOK)
}
