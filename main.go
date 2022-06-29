package main

import (
	"ApiRestWithGinGo/database"
	"ApiRestWithGinGo/models"
	"ApiRestWithGinGo/routes"
)

func main() {
	database.ConnectBd()
	models.Alunos = []models.Aluno{
		{
			Nome: "Roger",
			CPF:  "049.344.214-56",
			RG:   "58831",
		},
		{
			Nome: "Mesquita",
			CPF:  "049.344.214-56",
			RG:   "588231",
		},
	}

	routes.HandleRequests()

}
