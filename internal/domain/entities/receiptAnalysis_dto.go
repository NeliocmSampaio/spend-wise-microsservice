package entities

// ReceiptAnalysis holds the analyzed data returned from ChatGPT.
type ReceiptAnalysis struct {
	Date     string `json:"date"`
	Value    string `json:"value"`
	Category string `json:"category"`
}
