package repository

type ChatGPTClient struct{}

func (c *ChatGPTClient) AnalyzeReceipt(receiptText string) (string, error) {
	// Finge a chamada para a API do ChatGPT
	// Na prática, aqui seria feita a chamada para a API
	return `{
        "date": "2023-01-01",
        "amount": "100.00",
        "description": "Compra de Mercado",
        "category": "Alimentos"
    }`, nil
}
