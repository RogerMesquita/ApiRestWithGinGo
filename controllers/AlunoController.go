package controllers

import (
	"ApiRestWithGinGo/database"
	"ApiRestWithGinGo/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ListAll(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func FindId(c *gin.Context) {
	id := c.Params.ByName("id")
	aluno := models.Aluno{}
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado",
		})
		return
	}

	c.JSON(200, aluno)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno nao encontrado",
		})
		return
	}

	c.JSON(200, aluno)
}

func Edit(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := models.ValidaDadosAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	database.DB.Updates(&aluno)

	c.JSON(200, aluno)
}

func GetCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("id")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	c.JSON(200, aluno)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("id")

	c.JSON(200, gin.H{
		"Api Diz: ": "Eai" + nome + ",tudo belexa?",
	})
}
