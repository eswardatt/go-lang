package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func AddFeature(c echo.Context) error {
	feature := entities.Feature{}
	json.NewDecoder(c.Request().Body).Decode(&feature)
	servService.AddFeature(feature)

	return c.JSON(http.StatusOK, " ' Feature ' Added Successfully")
}
func GetFeatures(c echo.Context) error {
	featuresList := servService.GetFeatures()
	return c.JSON(http.StatusOK, featuresList)
}
func UpdateFeature(c echo.Context) error {
	id := c.Param("id")
	feature := entities.Feature{}
	json.NewDecoder(c.Request().Body).Decode(&feature)
	servService.UpdateFeature(feature, id)

	return c.JSON(http.StatusOK, " ' Feature ' Updated Successfully")
}
