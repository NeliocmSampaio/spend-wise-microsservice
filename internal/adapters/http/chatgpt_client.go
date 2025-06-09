package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"spendwise-microservice/internal/adapters/http/presenters"
	"spendwise-microservice/internal/domain/entities"
	"strings"
)

const (
	transactionPrompt = `Analise o seguinte texto de transação e identifique a categoria da transação. Por favor, responda apenas com o nome da categoria.
						Texto da transação: "[REPLACE_RECEIPT_TEXT]"
						Retorne as informacoes no seguinte formado:
						{
							"date": "01/01/2006",
							"value": "100.00",
							"category": "Alimentacao"
						}
						Categorias: [[REPLACE_CATEGORY]]
						Caso a categoria não encaixe nas informadas, retorne a nova categoria em uma única palavra.
						`
)

type ChatGPTClient struct {
	ApiKey string
}

func NewChtGPTClient(apiKey string) ChatGPTClient {
	return ChatGPTClient{
		ApiKey: apiKey,
	}
}

// AnalyzeReceipt interacts with OpenAI's API to analyze the receipt text.
func (c *ChatGPTClient) AnalyzeReceipt(receipt entities.Receipt) (entities.ReceiptAnalysis, error) {
	url := "https://api.openai.com/v1/chat/completions"
	method := "POST"

	content := strings.Replace(transactionPrompt, "[REPLACE_RECEIPT_TEXT]", receipt.ReceiptText, 1)
	content = strings.Replace(content, "[REPLACE_CATEGORY]", strings.Join(receipt.Categories, ", "), 1)
	// Prepare the JSON payload for the API request.
	payload := map[string]interface{}{
		"model": "gpt-4o-mini",
		"store": true,
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": content,
			},
		},
	}

	// Serialize payload to JSON.
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return entities.ReceiptAnalysis{}, fmt.Errorf("error encoding json: %v", err)
	}

	// Create a new HTTP request with the given payload.
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return entities.ReceiptAnalysis{}, fmt.Errorf("error creating request: %v", err)
	}

	// Set necessary headers for the request.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	// Initialize HTTP client and execute the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return entities.ReceiptAnalysis{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Decode the response body into a map.
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return entities.ReceiptAnalysis{}, fmt.Errorf("error decoding response: %v", err)
	}

	// Extract necessary information from the response.
	// Adapt this block to accommodate the actual structure.
	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	var analysis presenters.ReceiptAnalysis
	err = json.Unmarshal([]byte(message), &analysis)
	if err != nil {
		return entities.ReceiptAnalysis{}, fmt.Errorf("error reading message: %v", err)
	}

	return analysis.ToReceiptAnalysisEntity(), nil
}
