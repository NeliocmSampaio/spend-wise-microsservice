package handler

import (
	"net/http"
	"spendwise-microservice/internal/application"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	chatGPTClient := application.TransactionService{ChatGPTClient: repository.ChatGPTClient{}}
	r.POST("/analyze-receipt", func(c *gin.Context) {
		var input struct {
			Text string `json:"text"`
		}
		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		transaction, err := chatGPTClient.AnalyzeText(input.Text)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transaction)
	})
}
