package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"spendwise-microservice/internal/adapters/entrypoints/v1/presenters"
	"spendwise-microservice/internal/domain/ports"
)

type TransactionController struct {
	chatGPTClient ports.AIMotorPort
}

func NewTransactionController(chatGptClient ports.AIMotorPort) TransactionController {
	return TransactionController{
		chatGPTClient: chatGptClient,
	}
}

func (c *TransactionController) HandleTransaction(context *gin.Context) {
	var input presenters.ReceiptDto
	if err := context.BindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	transaction, err := c.chatGPTClient.AnalyzeReceipt(input.ToReceiptEntity())
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, transaction)
}
