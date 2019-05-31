package routes

import ("github.com/stretchr/testify/assert"
         "testing"
			"net/http" 
			 "net/http/httptest"
			"github.com/gin-gonic/gin")



	func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
				req, _ := http.NewRequest(method, path, nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				return w
			 }

func TestStartRouter(t *testing.T) {
	// Assert we encoded correctly,
   // the request gives a 200
    // Perform a GET request with that handler.
	router := gin.Default()

	w := performRequest(router, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
	w  = performRequest(router, "POST", "/api/v1/signin")
	assert.Equal(t, http.StatusOK, w.Code)
	w  = performRequest(router, "POST", "/api/v1/signup")
	assert.Equal(t, http.StatusOK, w.Code)

  

	}