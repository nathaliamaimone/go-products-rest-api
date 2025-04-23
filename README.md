# API de Produtos em Go

API REST desenvolvida em Go para gerenciamento de produtos com autenticação JWT.

## Tecnologias
- Go 1.21
- Framework Gin
- PostgreSQL
- Docker
- JWT Authentication

## Pré-requisitos
- Go 1.21+
- Docker & Docker Compose
- PostgreSQL

## Configuração

1. Crie um arquivo `.env` na raiz do projeto:
```env
JWT_SECRET_KEY=your-secret-key-here
JWT_EXPIRATION_HOURS=24
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password
DB_NAME=your-database

## Executando o Projeto

Para executar o projeto, siga os passos abaixo:

1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/go-products-rest-api.git
   ```

1. Instalando as dependências:
   ```bash
   go mod download
   ```

3. Inicie o banco de dados:
   ```bash
   docker compose up -d
   ```

4. Execute a aplicação:
   ```bash
   cd cmd && go run main.go
   ```

## API Endpoints

### Autenticação
| Método | Endpoint | Descrição | Acesso |
|--------|----------|-----------|---------|
| POST | `/register` | Registrar novo usuário | Público |
| POST | `/login` | Login e obter token JWT | Público |

### Produtos
| Método | Endpoint | Descrição | Acesso |
|--------|----------|-----------|---------|
| GET | `/products` | Listar todos os produtos | Público |
| GET | `/products/:id` | Buscar produto por ID | Público |
| POST | `/products` | Criar novo produto | Protegido |
| PUT | `/products/:id` | Atualizar produto | Protegido |
| PATCH | `/products/:id` | Atualizar produto parcialmente | Protegido |
| DELETE | `/products/:id` | Deletar produto | Protegido |

## Autenticação

Para acessar as rotas protegidas da API, é necessário realizar autenticação utilizando JSON Web Token (JWT). O token deve ser obtido através do endpoint de login e incluído no cabeçalho de todas as requisições protegidas.

### Como utilizar

1. Faça login através do endpoint `/login` para obter o token JWT
2. Inclua o token no cabeçalho Authorization usando o formato Bearer

