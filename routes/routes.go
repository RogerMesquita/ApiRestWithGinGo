package routes

import (
	"ApiRestWithGinGo/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ListAll)
	r.GET("/alunos/:id/details", controllers.Saudacao)
	r.POST("/alunos/new", controllers.New)
	r.GET("/alunos/:id/view", controllers.FindId)
	r.DELETE("/alunos/:id/delete", controllers.Delete)
	r.POST("/alunos/:id/edit", controllers.Edit)
	r.GET("/alunos/:id/cpf", controllers.GetCpf)

	r.Run()
}
