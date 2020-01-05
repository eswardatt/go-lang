package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

//getServe:returns all Services
func GetServ(c echo.Context) error {
	serviceList := servService.Get()
	return c.JSON(http.StatusOK, serviceList)
}
func AddService(c echo.Context) error {
	service := entities.Service{}
	json.NewDecoder(c.Request().Body).Decode(&service)
	servService.AddService(service)
	return c.JSON(http.StatusOK, "Service Added Successfully")
}



