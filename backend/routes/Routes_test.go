package routes

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequestLogin(r http.Handler, method, path string, t *testing.T) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader("{\"email\": \"ruiblaese@gmail.com\",\"password\": \"1234\"}"))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	return w
}

func performRequestWithToken(r http.Handler, method, path string, t *testing.T) *httptest.ResponseRecorder {
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InJ1aWJsYWVzZUBnbWFpbC5jb20iLCJleHAiOjE1NTk0MjQxMjIsIm9yaWdfaWF0IjoxNTU5NDIwNTIyfQ.kdIRkLjRc63VQvDcHECId45_8rlCr8QlAmVBcEG2tlE"
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader("{\"email\": \"ruiblaese@gmail.com\",\"password\": \"1234\"}"))
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	r.ServeHTTP(w, req)

	//assert.Equal(t, w.Code, http.StatusOK)
	return w
}

func TestStartRouter(t *testing.T) {
	/*
		os.Setenv("DB_HOST", "172.17.0.2")
		os.Setenv("DB_USER", "postgres")
		os.Setenv("DB_PASSWORD", "12345")
		os.Setenv("DB_DATABASE", "codenation")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("JWT_SECRET", "SecretJWTKeyCodeNation")

		fmt.Println(os.Getenv("DB_HOST"))
	*/

	ginRouter := gin.Default()
	router := StartRouter(ginRouter)

	w := performRequestLogin(router, "POST", "/api/v1/signin", t)
	assert.Equal(t, http.StatusOK, w.Code)
	/*
		w = performRequest(router, "GET", "/", t)
		assert.Equal(t, http.StatusOK, w.Code)
		w = performRequest(router, "POST", "/api/v1/signup", t)
		assert.Equal(t, http.StatusOK, w.Code)
		w = performRequest(router, "GET", "/user", t)
		assert.Equal(t, http.StatusOK, w.Code)
		w = performRequest(router, "GET", "/user/id", t)
		assert.Equal(t, http.StatusOK, w.Code)
		w = performRequest(router, "GET", "/customer", t)
		assert.Equal(t, http.StatusOK, w.Code)
		w = performRequest(router, "GET", "/customer/id", t)
		assert.Equal(t, http.StatusOK, w.Code)
	*/
}
