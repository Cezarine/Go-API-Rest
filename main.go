package main

import (
	"Go-API-Rest/models"
	"Go-API-Rest/routes"
	"fmt"
)

const porta = ":8000"

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"}}

	fmt.Println("Servidor rodando na porta", porta)
	routes.HandleRequest(porta)
}
