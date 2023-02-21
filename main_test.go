package main

import (
	"fmt"
	"gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetWelcomeMessage(t *testing.T) {
	mockResponse := `{"message":"Welcome bicho!"}`

	fmt.Println(mockResponse)
	r := SetUpRouter()
	r.GET("/", controllers.GetWelcomeMessage())
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}
