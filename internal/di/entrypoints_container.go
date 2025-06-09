package di

import "spendwise-microservice/internal/adapters/entrypoints/v1/controllers"

type EntrypointsContainer struct {
	TransactionController controllers.TransactionController
}

func BuildEntrypointsContainer(domainContainer DomainContainer) EntrypointsContainer {
	transactionController := controllers.NewTransactionController(&domainContainer.ChatGPTAdapter)

	return EntrypointsContainer{
		TransactionController: transactionController,
	}
}
