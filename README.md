# Boilerplate Clean Go

Este é um boilerplate para construção de aplicações em Go, com foco em uma estrutura limpa e bem organizada. O projeto é modular, com separação clara entre as diferentes camadas e responsabilidades.

## Estrutura do Projeto

- **cmd**: Contém a entrada principal e os testes da aplicação.
- **configs**: Armazena as configurações da aplicação.
- **internal**: Contém a lógica de negócios e a infraestrutura do sistema.
- **pkg**: Pacotes reutilizáveis em todo o projeto.
- **README.md**: Este arquivo.

## Como Usar

1. Clone o repositório.
2. Instale as dependências:
    ```bash
    go mod tidy
    ```
3. Execute a aplicação:
    ```bash
    go run cmd/main.go
    ```

## Contribuições

Contribuições para melhorar a estrutura ou adicionar novas funcionalidades são bem-vindas. Certifique-se de seguir as diretrizes de estilo e realizar testes adequados.
