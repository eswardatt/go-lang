package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetCountries(c echo.Context) error {
	return c.JSON(http.StatusOK, cntryService.GetCountries())
}

func GetStates(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, cntryService.GetStates(id))
};

func GetCities(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, cntryService.GetCities(id))
}
