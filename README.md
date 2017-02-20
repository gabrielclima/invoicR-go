# API REST com Golang
Uma API REST para invoices (notas fiscais) utilizando a linguagem Golang.

Usando as seguintes configurações
* Golang como linguagem
* Framework [gorila/mux]([https://github.com/gorilla/mux]) para controle de rotas e requisições (HTTP)
* MySQL como banco de dados

# Instruções

## Configuração banco de dados
Dentro da pasta `database` existe um arquivo nomeado como `invoices.sql`.
Você pode copiar todo o script `.sql` dentro do arquivo e rodar diretamente no banco de dados.

```
$ cd go_rest_api/database
$ mysql -u user -p password
$ source invoices.sql
```

## Rodando o programa
Com o banco de dados criado e configurado. Agora podemos rodar a API.
Podemos usar dois métodos para fazer isso:

1. Buildando
```
$ go build
```
Quando este comando é executado na raiz do projeto, ou seja, `go_rest_api/` um executável será
gerado na raiz do projeto. Se seu sistema for Unix/Linux, o nome será `go_rest_api`. Se seu sistema
for Windows, o nome será `go_rest_api.exe`.

2. Rodando diretamente
```
$ go run main.go
```
Esse comando faz o projeto ser compilado e executado.
