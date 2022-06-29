package routes

import (
	"ApiRestWithGinGo/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos/new", controllers.New)
	r.GET("/:nome", controllers.Saudacao)

	r.Run()
}
