# Estágio de construção
FROM golang:alpine AS builder

# Defina o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o código-fonte
COPY . .

# Copiar go mod e sum para o cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Execute os testes (opcional)
RUN go test ./cmd/server ./internal/infra -v

# Compile o aplicativo com flags para reduzir o tamanho do binário
RUN go build -ldflags="-s -w" -o app cmd/server/main.go

# Estágio de produção
FROM alpine:latest

# Recebe a versão como argumento de build
ARG VERSION=0.0.0

# Define a versão como variável de ambiente
ENV BUILD_VERSION=$VERSION

# Verifica o valor da versão
RUN echo "Build Version: $BUILD_VERSION"

# Copie o arquivo de aplicação compilado do estágio de construção
COPY --from=builder /app/app /app/app

# Copie os arquivos de configuração necessários para o container
COPY configs/.env /app/configs/.env

# Defina o diretório de trabalho para o aplicativo
WORKDIR /app

# Copie o arquivo start.sh para o contêiner e conceda permissões de execução
COPY deploy/scripts/start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Defina o comando padrão para iniciar o aplicativo (se o arquivo for start.sh)
CMD ["/app/start.sh"]
