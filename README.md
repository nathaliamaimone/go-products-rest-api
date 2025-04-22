# API de Produtos em Go

API REST desenvolvida em Go para gerenciamento de produtos.

## Tecnologias
- Go 1.21
- Framework Gin
- PostgreSQL
- Docker

## Pré-requisitos
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL

## Executando o Projeto

1. Inicie o banco de dados:

   ```bash
   docker compose up -d

2. Execute a aplicação:

   ```bash
   cd cmd && go run main.go

## Endpoints da API
- GET /products - Lista todos os produtos
- GET /products/:id - Busca produto por ID
- POST /products - Cria novo produto
- PUT /products/:id - Atualiza produto
- DELETE /products/:id - Remove produto
