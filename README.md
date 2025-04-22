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

Para executar o projeto, siga os passos abaixo:

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/go-products-rest-api.git
   ```

2. Inicie o banco de dados:
   ```bash
   docker compose up -d
   ```

3. Execute a aplicação:
   ```bash
   cd cmd && go run main.go
   ```

## Endpoints da API

### Produtos

| Método | Rota | Descrição |
|--------|------|-----------|
| GET | `/products` | Lista todos os produtos |
| GET | `/products/:id` | Busca produto por ID |
| POST | `/products` | Cria novo produto |
| PUT | `/products/:id` | Atualiza produto |
| PATCH | `/products/:id` | Atualiza produto parcialmente |
| DELETE | `/products/:id` | Remove produto |

