//go:build wireinject
// +build wireinject

package main

import (
	"pegawaiServices/app"

	"pegawaiServices/controller"
	"pegawaiServices/repository"
	"pegawaiServices/service"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var pegawaiSet = wire.NewSet(
	repository.NewPegawaiRepositoryImpl,
	wire.Bind(new(repository.PegawaiRepository), new(*repository.PegawaiRepositoryImpl)),
	service.NewPegawaiServiceImpl,
	wire.Bind(new(service.PegawaiService), new(*service.PegawaiServiceImpl)),
	controller.NewPegawaiControllerImpl,
	wire.Bind(new(controller.PegawaiController), new(*controller.PegawaiControllerImpl)),
)

func InitializedServer() *echo.Echo {
	wire.Build(
		app.GetConnection,
		wire.Value([]validator.Option{}),
		validator.New,
		pegawaiSet,
		app.NewRouter,
	)
	return nil
}
