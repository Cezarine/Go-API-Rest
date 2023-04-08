package main

import (
	"Go-API-Rest/database"
	"Go-API-Rest/routes"
	"fmt"
)

const porta = ":8000"

func main() {
	database.ConnectDataBase()
	fmt.Println("Servidor rodando na porta", porta)
	routes.HandleRequest(porta)
}
