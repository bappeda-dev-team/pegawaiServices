# API E-KAK 


### Kebutuhan

- Go versi 1.23.0
- MySQL versi 8.0.33
- [golang-migrate](https://github.com/golang-migrate/migrate)


### Cara install

- buat database bernama db_pegawai

- install golang-migrate di cmd

```sh
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- add create migration
``` sh
migrate create -ext sql -dir db create_new_table
```

- migrasi database

```sh
migrate -path db -database "mysql://root@tcp(localhost:3306)/db_pegawai" up
```

- add echo framework
```sh
go get github.com/labstack/echo/v4
```

```sh
go get github.com/labstack/echo/v4/middleware
```


### Run server

ketikkan perintah:

```sh
go run main.go
```

untuk hot reload(only macOS/linux)

install
```sh
go install github.com/cespare/reflex@latest
```

```sh
reflex -s -r '\.go$$' go run .
```


untuk menghentikan server, tekan Ctrl + c



### Documentasi API With Swagger Echo
- Untuk menampilkan Ui Swagger
```sh
http://localhost:8001/swagger/index.html
```

- Untuk install swagger echo
```sh
go install github.com/swaggo/swag/cmd/swag@latest
```
```sh
go get -d github.com/swaggo/swag/cmd/swag
```
```sh
go get -u github.com/swaggo/echo-swagger
```

