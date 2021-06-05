<h1 align="center">Api GoLang</h1>

## 📝 **Sobre**
Exemplo de Api utilizando a linguagem Go utilizando Gin como Router e Gorm como ORM.

## 🔨 **Tecnologias Utilizadas**

- [Go](https://golang.org/doc/)
- [Gin](https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [Docker-Compose](https://docs.docker.com/compose/)

## 📚 **Requisitos para o ambiente**
- Go 1.16.4
- Docker

## 🚀 **Começando**

### Clone o projeto do repositório com o comando abaixo
```sh
$ git clone https://github.com/repRSilva/api-golang.git
```

### Execute o banco de dados
```sh
$ docker-compose up -d
```

### Executar aplicação em modo produção
```sh
$ go run main.go
```

## ⚙️ **Obtendo os resultados**

### Lista todos os usuários cadastrados
```
MÉTODO GET
http://localhost:5000/api/v1/users
```

### Lista um usuário cadastrado pelo seu id
```
MÉTODO GET
http://localhost:5000/api/v1/users/"idUsuario"
```

### Criar um novo usuário
```
MÉTODO POST
http://localhost:5000/api/v1/users

- body
{
  "name": "String",
  "age": Integer,
}
```

### Atualizar um usuário
```
MÉTODO PUT
http://localhost:5000/api/v1/users

- body
{
  "id":Integer,
  "name": "String",
  "age": Integer,
}
```

### Deletar um usuário
```
MÉTODO DELETE
http://localhost:5000/api/v1/users/"idUsuario"
```

Desenvolvido por [Rafael Silva](https://github.com/repRSilva/) ;D 🚀
