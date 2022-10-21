# Scalable Web Services Golang - Hacktiv8

## Useful Library

```
// Gin gonic (Framework for Rest API)
go get -u github.com/gin-gonic/gin

// Postgres Driver
go get -u github.com/lib/pq

// Gorm Postgres Driver
go get gorm.io/driver/postgres

// Gorm - ORM Library
go get -u gorm.io/gorm
go get -u "github.com/jinzhu/gorm/dialects/mysql"

// Swaggo - Documentation
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template

// Gorilla Mux - Gin gonic alternative
go get -u github.com/gorilla/mux

// JWT
go get -u github.com/dgrijalva/jwt-go

// Cryptography
go get golang.org/x/crypto

// Testing
go get github.com/stretchr/testify


// Additional notes
-u = get latest package
```

## Basic Go Command

```
// Lihat versi
go version

// Inisialisasi go modules
go mod init <nama-project>

// Mengeksekusi file go
go run main.go

// Mengeksekusi fule go yang mengimport package lain di satu direktori
go run main.go library.go

// Unit test
go test main_test.go

// Kompilasi file menjadi exe
go build -o program.exe

// Mendownload package
go get github.com/segmentio/kafka-go

// Memvalidasi dependensi yang belum terdonlot
go mod tidy

// Vendoring (advance command)
go mod vendor
-u = get latest package
```
