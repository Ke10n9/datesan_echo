package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"app/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, Echo World!!")
	})
	
	e.POST("/signup", controller.Signup) // POST /signup
	e.POST("/login", controller.Login) // POST /login

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(controller.Config)) // /api 下はJWTの認証が必要
	api.POST("/menu/create", controller.AddMenu)
	api.GET("/menu/:id", controller.GetMenu)
	api.DELETE("/menu/:id", controller.DeleteMenu)
	api.PUT("/menu/:id", controller.UpdateMenu)

	e.Logger.Fatal(e.Start(":8080"))
}