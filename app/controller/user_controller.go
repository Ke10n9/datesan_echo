package controller

import (
	"net/http"

	"github.com/labstack/echo"
  "app/model"
)

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
			return err
	}

	if user.Email == "" || user.Password == "" {
			return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid name or password",
			}
	}

	if u := model.FindUserByEmail(user.Email); u.ID != 0  {
			return &echo.HTTPError{
					Code:    http.StatusConflict,
					Message: "email already exists",
			}
	}

	model.CreateUser(user)
	user.Password = ""

	return c.JSON(http.StatusCreated, user)
}