# Scalable Web Services Golang - Hacktiv8

## Useful Library
```
// Gin gonic (Framework for Rest API)
go get -u github.com/gin-gonic/gin

// SQL Driver
go get -u github.com/lib/pq

// Gorm - ORM Library
go get -u gorm.io/gorm 

// Swagger - Documentation
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/swag/http-swagger
go get -u github.com/alecthomas/template


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
