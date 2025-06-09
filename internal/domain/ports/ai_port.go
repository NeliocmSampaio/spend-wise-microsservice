package ports

import "spendwise-microservice/internal/domain/entities"

type AIMotorPort interface {
	AnalyzeReceipt(receipt entities.Receipt) (entities.ReceiptAnalysis, error)
}
