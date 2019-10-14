package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/AashrayAnand/tripit/session"
	"github.com/AashrayAnand/tripit/util"
	"github.com/gin-gonic/gin"
)

// server to use for tests
var router = util.RunServer()

// utility function which serves HTTP request, and passes response
// to a function parameter, which checks results, used by all tests,
// credit: https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
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
