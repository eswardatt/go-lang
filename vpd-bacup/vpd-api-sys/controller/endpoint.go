package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func Getendpoints(c echo.Context) error {
	endpoints := servService.GetEndpoints()
	return c.JSON(http.StatusOK, endpoints)
}

func AddEndpoint(c echo.Context) error {
	endpoint := entities.EndPoint{}
	json.NewDecoder(c.Request().Body).Decode(&endpoint)
	servService.AddEndpoint(endpoint)

	return c.JSON(http.StatusOK, " ' Endpoint ' Added Successfully")
}
