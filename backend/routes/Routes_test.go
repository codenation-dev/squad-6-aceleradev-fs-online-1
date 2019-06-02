package routes

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string, t *testing.T) *httptest.ResponseRecorder {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJ1aWJsYWVzZUBnbWFpbC5jb20iLCJleHAiOjE1NTk0MjQxMjIsIm9yaWdfaWF0IjoxNTU5NDIwNTIyfQ.kdIRkLjRc63VQvDcHECId45_8rlCr8QlAmVBcEG2tlE"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	return w
}

func TestStartRouter(t *testing.T) {
	// Assert we encoded correctly,
	// the request gives a 200
	// Perform a GET request with that handler.
	router := StartRouter()

	w := performRequest(router, "GET", "/",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "POST", "/api/v1/signin",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "POST", "/api/v1/signup",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/user",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/user/id",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/customer",t)
	assert.Equal(t, http.StatusOK, w.Code)
	w = performRequest(router, "GET", "/customer/id",t)
	assert.Equal(t, http.StatusOK, w.Code)

}
