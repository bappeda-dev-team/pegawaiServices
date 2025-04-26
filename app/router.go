package app

import (
	"pegawaiServices/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(pegawaiController controller.PegawaiController) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/pegawai/create", pegawaiController.Create)
	e.PUT("/pegawai/update/:id", pegawaiController.Update)
	e.DELETE("/pegawai/delete/:id", pegawaiController.Delete)
	e.GET("/pegawai/detail/:id", pegawaiController.FindById)
	e.GET("/pegawai/findall", pegawaiController.FindAll)
	e.GET("/pegawai/findbynip/:nip", pegawaiController.FindByNip)

	return e
}
