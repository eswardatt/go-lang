package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func AddMenu(c echo.Context) error {
	menu := entities.Menu{}
	json.NewDecoder(c.Request().Body).Decode(&menu)
	servService.AddMenu(menu)

	return c.JSON(http.StatusOK, " ' Menu ' Added Successfully")
}
func GetMenu(c echo.Context) error {
	menuList := servService.GetMenu()
	return c.JSON(http.StatusOK, menuList)
}