package presenters

import "spendwise-microservice/internal/domain/entities"

// ReceiptAnalysis holds the analyzed data returned from ChatGPT.
type ReceiptAnalysis struct {
	Date     string `json:"date"`
	Value    string `json:"value"`
	Category string `json:"category"`
}

func (p *ReceiptAnalysis) ToReceiptAnalysisEntity() entities.ReceiptAnalysis {
	return entities.ReceiptAnalysis{
		Date:     p.Date,
		Value:    p.Value,
		Category: p.Category,
	}
}
