package application

import (
	"spendwise-microservice/internal/domain"
	"spendwise-microservice/internal/infrastructure/repository"
)

type TransactionService struct {
	ChatGPTClient repository.ChatGPTClient
}

func (ts *TransactionService) AnalyzeText(receiptText string) (*domain.Transaction, error) {
	response, err := ts.ChatGPTClient.AnalyzeReceipt(receiptText)
	if err != nil {
		return nil, err
	}

	transaction := parseResponseToTransaction(response)
	return transaction, nil
}

func parseResponseToTransaction(response string) *domain.Transaction {
	// Simulação de parsing simples
	return &domain.Transaction{
		Date:        "2023-01-01",
		Amount:      "100.00",
		Description: "Compra de Mercado",
		Category:    "Alimentos",
	}
}
