package main

import (
	"fmt"
	"os"
	_ "pegawaiServices/docs"
	"pegawaiServices/helper"

	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) *echo.Echo {
	return e
}

// @title Pegawai Service API
// @version 1.0
// @description API For Master Pegawai Services
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8001
// @BasePath /
// @schemes http

func main() {
	server := InitializedServer()
	host := os.Getenv("host")
	port := os.Getenv("port")

	addr := fmt.Sprintf("%s:%s", host, port)
	//kedepannya gunakan jwt rs256
	// server.Use(middleware.AuthMiddleware)

	err := server.Start(addr)
	helper.PanicIfError(err)
}
