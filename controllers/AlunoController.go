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
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"Api Diz: ": "Eai" + nome + ",tudo belexa?",
	})
}
