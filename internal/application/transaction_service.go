package application

// import (
// 	repository "spendwise-microservice/internal/adapters/http"
// 	"spendwise-microservice/internal/domain/entities"
// )

// type TransactionService struct {
// 	ChatGPTClient repository.ChatGPTClient
// }

// func (ts *TransactionService) AnalyzeText(receiptText string) (*entities.Transaction, error) {
// 	response, err := ts.ChatGPTClient.AnalyzeReceipt(receiptText)
// 	if err != nil {
// 		return nil, err
// 	}

// 	transaction := parseResponseToTransaction(response)
// 	return transaction, nil
// }

// func parseResponseToTransaction(response string) *entities.Transaction {
// 	// Simulação de parsing simples
// 	return &entities.Transaction{
// 		Date:        "2023-01-01",
// 		Amount:      "100.00",
// 		Description: "Compra de Mercado",
// 		Category:    "Alimentos",
// 	}
// }
