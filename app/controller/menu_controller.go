package controller

import (
	"net/http"
	// "fmt"
	"strconv"

	"github.com/labstack/echo"
	"app/model"
)

func AddMenu(c echo.Context) error {
	menu := new(model.Menu)
	if err := c.Bind(menu); err != nil {
			return err
	}

	if menu.Time_id == 0 {
			return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: "invalid date or time",
			}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
        return echo.ErrNotFound
	}

	menu.User_ID = uid
	model.CreateMenu(menu)

	return c.JSON(http.StatusCreated, menu)
}

func GetMenu(c echo.Context) error {
	menuId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	menu := model.FindMenuById(menuId)
	return c.JSON(http.StatusOK, menu)
}

func DeleteMenu(c echo.Context) error {
	menuID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			return echo.ErrNotFound
	}

	if err := model.DeleteMenu(&model.Menu{ID: menuID}); err != nil {
			return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateMenu(c echo.Context) error {
	menuID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	time_id, _ := strconv.Atoi(c.FormValue("time_id"))

	menu := model.FindMenuById(menuID)
	// if len(menus) == 0 {
	// 		return echo.ErrNotFound
	// }
	
	menu.Time_id = time_id
	menu.Date = c.FormValue("date")

	model.DB.Save(&menu)

	return c.NoContent(http.StatusNoContent)
}