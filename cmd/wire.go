//go:build wireinject
// +build wireinject

package cmd

import (
	"golang-test/configs"
	"golang-test/handlers"
	"golang-test/repositories"
	"golang-test/routes"
	"golang-test/services"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

func InitializeRouter() *echo.Echo {
	wire.Build(
		configs.NewEnv,
		configs.NewDB,
		configs.NewRedis,
		repositories.NewRedisRepository,
		repositories.NewProductRepository,
		services.NewProductService,
		handlers.NewProductHandler,
		routes.NewRouter,
	)
	return &echo.Echo{}
}
