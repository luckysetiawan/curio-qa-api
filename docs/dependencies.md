# Curio QA API

## Dependencies

### Go mod

Go mod initialization
```
go mod init github.com/luckysetiawan/curio-qa-api
```
When adding and using packages do
```
go mod tidy
```

## Packages

Commands used to get the required packages
```
go get github.com/joho/godotenv
go get github.com/gorilla/mux
go get github.com/rs/cors
go get go.mongodb.org/mongo-driver/mongo
go get github.com/redis/go-redis/v9
go get golang.org/x/crypto
go get github.com/golang-jwt/jwt/v5
go get github.com/stretchr/testify
```
Packages docs:
1. godotenv: [docs](https://github.com/joho/godotenv)
1. gorilla/mux: [docs](https://github.com/gorilla/mux)
1. rs/cors: [docs](https://github.com/rs/cors)
1. mongo-driver/mongo: [docs](https://github.com/mongodb/mongo-go-driver)
1. go-redis/v9: [docs](https://github.com/redis/go-redis)
1. x/crypto: [docs](https://pkg.go.dev/golang.org/x/crypto)
1. jwt/v5: [docs](https://github.com/golang-jwt/jwt)
1. stretchr/testify: [docs](https://github.com/stretchr/testify)
