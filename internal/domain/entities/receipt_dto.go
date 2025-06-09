package entities

type Receipt struct {
	ReceiptText string   `json:"receipt-text"`
	Categories  []string `json:"categories"`
}
