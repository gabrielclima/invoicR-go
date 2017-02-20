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

A API está configurada para rodar na porta `8080`

## Rotas

### GET /invoices

Endpoint usado para retornar todas as notas fiscais ativas no banco de dados.

```
$ curl -H "Content-Type: application/json" -H "Authorization: token#app1" localhost/invoices 
```

Resposta esperada 
```
Status code 200

[
  {
    "id": "1",
    "document": "0",
    "description": "Uma nota fiscal qualquer",
    "amount": "0",
    "reference_mounth": "0",
    "reference_year": "0",
    "created_at": "2017-02-03T21:27:28Z",
    "is_active": "1"
  }
]
```

### GET /invoices{documento}

Endpoint usado para retornar uma nota fiscal pelo número do documento

```
$ curl -H "Content-Type: application/json" -H "Authorization: token#app1"  localhost/invoices/{documento} 
```

Resposta esperada 
```
Status code 200

[
  {
    "id": "1",
    "document": "0",
    "description": "Uma nota fiscal qualquer",
    "amount": "0",
    "reference_mounth": "0",
    "reference_year": "0",
    "created_at": "2017-02-03T21:27:28Z",
    "is_active": "1"
  }
]
```

### POST /invoices

Endpoint usado para retornar uma nota fiscal pelo número do documento

```
$ curl -H "Content-Type: application/json" -H "Authorization: token#app1" -X POST -d
{
    "document":"1222",
    "description":"Uma nota fiscal qualquer",
    "amount": "123.00",
    "reference_mounth":"12",
    "reference_year":"2014"
} localhost:8080/invoices
```

Resposta esperada 
```
Status code 201

[
  {
    "id": "1",
    "document": "0",
    "description": "Uma nota fiscal qualquer",
    "amount": "0",
    "reference_mounth": "0",
    "reference_year": "0",
    "created_at": "2017-02-03T21:27:28Z",
    "is_active": "1"
  }
]
```

### DELETE /invoices{documento}

Endpoint usado para fazer um soft delete, ou seja, inativar uma nota fiscal pelo número do documento

```
$ curl -H "Content-Type: application/json" -H "Authorization: token#app1" -X DELETE -d localhost:8080/invoices/{document}
```

Resposta esperada 
```
Status code 200
```





