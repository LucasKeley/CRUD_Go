# CRUD_Go

CRUD_Go é uma API escrita em Go para gerenciar usuários, implementando as operações de cadastro, leitura, atualização e exclusão (CRUD), além do fluxo de autenticação com JWT. O projeto serve como laboratório para boas práticas no ecossistema Go com Gin, MongoDB e testes automatizados.

## Principais recursos
- Endpoints REST para criação, consulta, atualização, exclusão e autenticação de usuários.
- Persistência em MongoDB utilizando o driver oficial (`mongo-driver`).
- Validações automáticas com `validator.v10` integradas ao Gin.
- Geração e validação de tokens JWT para proteger rotas sensíveis.
- Cobertura de testes com `testing`, `testify`, `gomock` e um ambiente MongoDB efêmero para integração.
- Observabilidade básica via `zap` para logs estruturados.

## Estrutura do projeto
```
├── Dockerfile
├── init_dependencies.go
├── main.go
├── src
│   ├── configuration        # Conexão com MongoDB, logger, validações e erros REST
│   ├── controller           # Camada HTTP (handlers, rotas e modelos da API)
│   ├── model                # Domínio da aplicação (serviços, repositórios e entidades)
│   └── tests                # Testes integrados end-to-end e utilidades de teste
└── README.md
```

## Pré-requisitos
- Go 1.24 ou superior (conforme `go.mod`).
- MongoDB 6.x+ acessível na URL definida em `MONGODB_URL`.
- Make (opcional) para automatizar comandos locais.
- Docker e Docker Compose (opcionais) para executar via contêiner.

## Configuração
1. Clone o repositório e instale as dependências:
   ```bash
   git clone https://github.com/LucasKeley/CRUD_Go.git
   cd CRUD_Go
   go mod download
   ```
2. Crie um arquivo `.env` com as variáveis necessárias. Um exemplo inicial:
   ```dotenv
   MONGODB_URL=mongodb://localhost:27017
   MONGODB_USER_DB=users
   JWT_SECRET_KEY=troque-por-uma-chave-secreta
   ```
   > Ajuste os valores para o seu ambiente. Durante os testes de integração é utilizado `dockertest`, portanto não é necessário subir o banco manualmente.

## Executando localmente
1. Inicie um MongoDB local se ainda não estiver rodando.
2. Execute a aplicação:
   ```bash
   go run ./...
   ```
3. A API ficará disponível em `http://localhost:8080`.

### Principais endpoints
| Método | Caminho                     | Descrição                                    | Autenticação |
| ------ | -------------------------- | -------------------------------------------- | ------------ |
| POST   | `/createUser`              | Cria um novo usuário                         | Não          |
| POST   | `/login`                   | Autentica o usuário e retorna um token JWT   | Não          |
| GET    | `/getUserById/:userId`     | Busca usuário pelo ID                        | Bearer JWT   |
| GET    | `/getUserByEmail/:userEmail` | Busca usuário pelo e-mail                  | Bearer JWT   |
| PUT    | `/updateUser/:userId`      | Atualiza dados de um usuário                 | Não*         |
| DELETE | `/deleteUser/:userId`      | Remove um usuário                            | Não*         |

\* No código atual as rotas de escrita ainda não exigem token JWT, mas isso pode ser ajustado facilmente na camada de rotas.

### Exemplo de chamada
```bash
curl -X POST http://localhost:8080/createUser \
  -H 'Content-Type: application/json' \
  -d '{
        "email": "user@example.com",
        "password": "Senha@123",
        "name": "User Example",
        "age": 30
      }'
```

## Testes
- Testes unitários e de integração:
  ```bash
  go test ./...
  ```
- Os testes da pasta `src/tests` inicializam um MongoDB temporário utilizando `dockertest`. Certifique-se de que o Docker esteja disponível ao executá-los.

## Docker
Construindo e executando a aplicação em contêiner:
```bash
docker build -t crud_go .
docker run --env-file .env -p 8080:8080 crud_go
```
> O `Dockerfile` usa build multi-stage; garanta que a imagem base esteja disponível para a versão do Go declarada.

## Contribuição
1. Abra uma issue descrevendo a melhoria ou correção.
2. Crie um fork, faça suas alterações em um branch dedicado e escreva testes cobrindo o novo comportamento.
3. Submeta um Pull Request detalhando o que foi alterado.

## Licença
Distribuído sob a licença MIT. Consulte o arquivo LICENSE (quando disponível) para mais informações.
