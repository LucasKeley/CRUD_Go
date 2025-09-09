# Use a versão mais recente do Go como imagem base para o estágio de compilação
FROM golang:1.25

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia os arquivos de código-fonte, dependências e módulos
COPY src src
COPY docs docs
COPY go.mod go.mod
COPY go.sum go.sum
COPY init_dependencies.go init_dependencies.go
COPY main.go main.go

# Compila a aplicação Go, criando um binário estático e pequeno
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
  go build -o CRUD_Go .

# Usa uma imagem Alpine mais leve para o estágio de execução final
FROM golang:1.25 as runner

# Adiciona um usuário não-root para segurança
RUN adduser -D lucaske

# Copia o binário compilado do estágio anterior
COPY --from=BUILDER /app/CRUD_Go /app/CRUD_Go

# Define as permissões de propriedade e execução
RUN chown -R lucaske:lucaske /app
RUN chmod +x /app/CRUD_Go

# Expõe a porta que a aplicação escuta
EXPOSE 8080

# Troca para o usuário não-root
USER lucaske

# Comando para rodar a aplicação quando o contêiner iniciar
CMD ["./CRUD_Go"]

