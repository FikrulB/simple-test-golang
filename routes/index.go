package routes

import (
	"github.com/labstack/echo/v4"
	"golang-test/handlers"
	middlewares "golang-test/middleware"
)

func NewRouter(productHandler handlers.ProductHandler) *echo.Echo {
	e := echo.New()
	e.Validator = middlewares.NewValidator()

	e.POST("/product", productHandler.CreateProduct)
	e.GET("/product", productHandler.GetProducts)

	return e
}
