# SpendWise Microserviço

## Introdução

O SpendWise é um microsserviço desenvolvido em Golang para automatizar a análise de textos de comprovantes financeiros utilizando a API do ChatGPT. Através da arquitetura hexagonal, o serviço recebe textos de comprovantes e retorna dados estruturados, como data, valor, descrição e categoria da transação. Este projeto visa aprimorar a precisão da extração de informações financeiras e facilitar o gerenciamento de finanças pessoais.

## Como Rodar o Microsserviço

### Pré-requisitos

- **Go**: Certifique-se de que a linguagem Go está instalada em seu sistema. Para instalação, acesse [https://golang.org/dl/](https://golang.org/dl/).
- **Docker**: Opcionalmente, você pode rodar o serviço usando Docker para facilitar o deployment e a consistência do ambiente. Instale-o a partir de [https://www.docker.com/](https://www.docker.com/).

### Passos para Rodar Localmente

'''
    go run cmd/service/main.go
'''