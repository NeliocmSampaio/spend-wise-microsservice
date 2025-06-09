package di

import (
	"spendwise-microservice/config"
	"spendwise-microservice/internal/adapters/http"
)

type DomainContainer struct {
	ChatGPTAdapter http.ChatGPTClient
}

func BuildDomainContainer(config config.Configuration) DomainContainer {
	chatGPTClient := http.NewChtGPTClient(config.API.Key)

	return DomainContainer{
		ChatGPTAdapter: chatGPTClient,
	}
}
