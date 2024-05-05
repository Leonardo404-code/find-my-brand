# Find my Brand

Find my brand é uma aplicação full-stack que realiza pesquisas no Google e verifica quais concorrentes estão se aproveitando dos seus termos de marca nos resultados patrocinados.

## Documentação da API

veja em [swagger.json](/docs/swagger.json)

## Como executar

### Download dependencias do projeto

```go
go mod vendor
```

### Execute o projeto

- Configure as variáveis de ambiente listadas abaixo
- No diretório raiz, execute:

```go
go run ./cmd/server/main.go
```

Você deve ver a seguinte mensagem:

![run output](/docs/images/run-output.png)

## Testes

- No diretório raiz, execute:

```go
go test ./... -v
```

Você deve ver a seguinte mensagem:

![test output](/docs/images/test.png)

## Variáveis de ambiente
Name | Description | Default
---  | ---         | ---
PORT | Port where the application listen for HTTP requests | `3000`
API_KEY | API key for authenticated requests | none
GOEMAIL_USER | user for send emails | none
GOEMAIL_PASSWORD | User password for send emails | none