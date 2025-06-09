package main

import (
	"fmt"
	"spendwise-microservice/config"
	"spendwise-microservice/internal/adapters/handler"
	"spendwise-microservice/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig("../../config/config.local.yaml")
	fmt.Println(config)

	r := gin.Default()

	domainContainer := di.BuildDomainContainer(config)
	entrypointsContainer := di.BuildEntrypointsContainer(domainContainer)

	handler.SetupRoutes(r, domainContainer, entrypointsContainer)
	r.Run(":8080")
}
