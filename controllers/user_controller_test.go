package controllers

import (
	"bytes"
	"fmt"
	"gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUp() *gin.Engine {
	router := gin.Default()
	configs.ConnectDB()
	return router
}

func TestGetWelcomeMessage(t *testing.T) {
	mockResponse := `{"message":"Welcome bicho!"}`
	r := SetUp()
	r.GET("/", GetWelcomeMessage())
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	defer w.Result().Body.Close()
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUser(t *testing.T) {
	jsonBody := []byte(`{"name": "Kantona", "location": "Tbilisi",
"title": "Football player"}`)
	bodyReader := bytes.NewReader(jsonBody)
	r := SetUp()
	r.POST("/users", CreateUser())
	req, _ := http.NewRequest(http.MethodPost, "/users", bodyReader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
	fmt.Println(string(responseData))

}
