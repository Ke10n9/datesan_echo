package router

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"app/controller"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, Echo World!!")
	})
	
	e.POST("/signup", controller.Signup) // POST /signup
	// e.POST("/login", controller.Login) // POST /login

	e.POST("/menu/create", controller.AddMenu)
	
	return e
}