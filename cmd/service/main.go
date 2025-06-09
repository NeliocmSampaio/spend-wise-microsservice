package main

import (
	"spendwise-microservice/internal/infrastructure/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handler.SetupRoutes(r)
	r.Run(":8080") // Inicia o servidor na porta 8080
}
