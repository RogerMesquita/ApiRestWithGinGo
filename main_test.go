package main

import (
	"ApiRestWithGinGo/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTestRoutes() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestStatusCodeVerification(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/alunos/:id/details", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/alunos/roger/details", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("Status error: valor recebido foi %d e o esperado é %d", response.Code, http.StatusOK)
	}

}

//utilizando testify
func TestAssert(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/alunos/:id/details", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/alunos/roginho/details", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Deveriam Ser Iguais")

	mockResponse := `{"Api Diz: ":"Eairoginho,tudo belexa?"}`

	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, mockResponse, string(responseBody))
}
