package presenters

import "spendwise-microservice/internal/domain/entities"

type ReceiptDto struct {
	ReceiptText string   `json:"receipt-text"`
	Categories  []string `json:"categories"`
}

func (d *ReceiptDto) ToReceiptEntity() entities.Receipt {
	return entities.Receipt{
		ReceiptText: d.ReceiptText,
		Categories:  d.Categories,
	}
}
