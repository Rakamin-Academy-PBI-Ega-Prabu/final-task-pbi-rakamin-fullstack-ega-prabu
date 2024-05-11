# Simple Photo Manager API

My first small project, as a beginner, building Backend API with Golang

## Depedencies

- [gin](https://github.com/gin-gonic/gin)
- [gorm](https://github.com/go-gorm/gorm)
- [jwt-go](https://github.com/golang-jwt/jwt)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [godotenv](https://github.com/lpernett/godotenv)
- [go-mysql driver](https://github.com/go-sql-driver/mysql)

## Extra Dependency for Local Development

- [CompileDaemon](https://github.com/githubnemo/CompileDaemon)

## Functionalities

1. User registration and login
2. User authorization using JWT (for accessing photo manager)
3. Basic photo manager (upload, edit, and delete)
4. User validation input

## Installation

### Prerequisite

1. Go installed, for more information how to install Go look up here [Go Download and Install](https://go.dev/doc/install).
2. MySQL Server, you can use local or remote server as long as you have the database url. For local server, you can use [XAMPP](https://www.apachefriends.org) or somethins similliar.

### Steps

First, clone the repository

```git
git clone https://github.com/feratyusa/final-task-pbi-rakamin-fullstack-prabu
```

Second, move into the folder repository. Create .env file by copying .env-example.

```bash
cp .env-example .env
```

This .env file is the server configuration of the application that define the port and database that will be used, and also secret string that will be used for JWT Signing. Below is an example of the .env file that you can use.

```env
PORT=3000
SECRET=random_string_sadqowe12319vas3c
DB_URL=user:pass@tcp(localhost:3306)/db_userapp?parseTime=true
```

Start building the application with:

```go
go build
```

A file will be created with .exe, you can start the app by running the application or in the terminal:

```bash
./userapp.exe
```

Localhost server will be running and you can test it using Postman or other similliar application. You can also install the go application and run it through terminal. For more information look on [Go Compile and Install](https://go.dev/doc/tutorial/compile-install) instructions.

## References

Thank you very much for all the people in the forum and people that providing tutorial where I can boost my learning of Golang. I also recommend these references if you are new to Golang.

- [Go Getting Started](https://go.dev/doc/tutorial/getting-started)
- [Gin Documentation](https://gin-gonic.com/docs/)
- [Gorm Documentation](https://gorm.io/index.html)
- [Coding with Roby - Go Tutorial](https://www.youtube.com/watch?v=-c0yEogFlvM&list=PL-LRDpVN2fZAluCzYNZdSCfJVQXe5ly90)
