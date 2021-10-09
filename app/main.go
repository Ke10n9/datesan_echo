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
	// e.POST("/login", controller.Login) // POST /login

	e.POST("/menu/create", controller.AddMenu)
	e.GET("/menu/:id", controller.GetMenu)
	e.DELETE("/menu/:id", controller.DeleteMenu)
	e.PUT("/menu/:id", controller.UpdateMenu)

	e.Logger.Fatal(e.Start(":8080"))
}