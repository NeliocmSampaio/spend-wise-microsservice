package handler

import (
	"spendwise-microservice/internal/di"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, domainContainer di.DomainContainer, entrypointsContainer di.EntrypointsContainer) {
	r.POST("/analyze-receipt", entrypointsContainer.TransactionController.HandleTransaction)
}
