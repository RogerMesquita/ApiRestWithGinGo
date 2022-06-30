package main

import (
	"ApiRestWithGinGo/controllers"
	"ApiRestWithGinGo/database"
	"ApiRestWithGinGo/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var ID int

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
func TestFindForId(t *testing.T) {
	database.ConnectBd()
	CreateAlunoMock()
	defer DeleteAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos/:id/view", controllers.FindId)
	path := "/alunos/" + strconv.Itoa(ID) + "/view"

	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	var alunoMock models.Aluno

	json.Unmarshal(response.Body.Bytes(), &alunoMock)

	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome)
}

func TestDeleteAluno(t *testing.T) {
	database.ConnectBd()
	CreateAlunoMock()
	r := SetupTestRoutes()
	r.DELETE("/alunos/:id/delete", controllers.Delete)
	path := "/alunos/" + strconv.Itoa(ID) + "/delete"
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditAluno(t *testing.T) {
	database.ConnectBd()
	CreateAlunoMock()
	r := SetupTestRoutes()
	defer DeleteAlunoMock()
	r.POST("/alunos/:id/edit", controllers.Edit)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste editado", CPF: "12345678903", RG: "123456710"}
	jsonValue, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID) + "/edit"
	req, _ := http.NewRequest("POST", path, bytes.NewBuffer(jsonValue))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, req) // faz a chamada de fato no server

	alunoMock := models.Aluno{}
	json.Unmarshal(response.Body.Bytes(), &alunoMock)
	assert.Equal(t, "12345678903", alunoMock.CPF)

}

func CreateAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeleteAlunoMock() {
	aluno := models.Aluno{}
	database.DB.Delete(&aluno, ID)
}
