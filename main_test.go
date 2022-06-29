package main

import (
	"ApiRestWithGinGo/controllers"
	"ApiRestWithGinGo/database"
	"ApiRestWithGinGo/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
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

func TestListAll(t *testing.T) {
	database.ConnectBd()
	CreateAlunoMock()

	defer DeleteAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos", controllers.ListAll)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestFindForCpf(t *testing.T) {
	database.ConnectBd()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos/:id/cpf", controllers.GetCpf)
	req, _ := http.NewRequest("GET", "/alunos/12345678901/cpf", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

var ID int

func CreateAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	aluno := models.Aluno{}
	database.DB.Delete(&aluno, ID)
}
