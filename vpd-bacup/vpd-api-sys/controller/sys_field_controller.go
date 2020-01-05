package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

//GetAllSysValues :
func GetAllSysValues(c echo.Context) error {
	return c.JSON(http.StatusOK, valueService.RetrieveAll())
}

//GetBySysValue :
func GetBySysValue(c echo.Context) error {
	sysValue := c.Param("sys-value-id")
	return c.JSON(http.StatusOK, valueService.RetrieveByValue(sysValue))
}

//AddSysValue :
func AddSysValue(c echo.Context) error {
	sysValue := entities.SysFieldValues{}
	json.NewDecoder(c.Request().Body).Decode(&sysValue)
	sysValueSave := valueService.Insert(sysValue)
	return c.JSON(http.StatusOK, sysValueSave)
}

//UpdateSysValue :
func UpdateSysValue(c echo.Context) error {
	sysValue := c.Param("sys-value-id")
	sysValueDetails := entities.SysFieldValues{}
	json.NewDecoder(c.Request().Body).Decode(&sysValue)
	return c.JSON(http.StatusOK, valueService.UpdateByValue(sysValue, sysValueDetails))
}

//DeleteSysValue :
func DeleteSysValue(c echo.Context) error {
	sysValue := c.Param("sys-value-id")
	return c.JSON(http.StatusOK, valueService.DeleteByValue(sysValue))
}

//GetAllSysCodes :
func GetAllSysCodes(c echo.Context) error {
	return c.JSON(http.StatusOK, codeService.RetrieveAll())
}

//GetBySysCode :
func GetBySysCode(c echo.Context) error {
	sysCode := c.Param("sys-code-id")
	return c.JSON(http.StatusOK, codeService.RetrieveByCode(sysCode))
}

//AddSysCode :
func AddSysCode(c echo.Context) error {
	sysCode := entities.SysFieldCode{}
	json.NewDecoder(c.Request().Body).Decode(&sysCode)
	sysCodeSave := codeService.Insert(sysCode)
	return c.JSON(http.StatusOK, sysCodeSave)
}

<<<<<<< HEAD

=======
//UpdateSysCode :
func UpdateSysCode(c echo.Context) error {
	sysCode := c.Param("sys-code-id")
	sysCodeDetails := entities.SysFieldCode{}
	json.NewDecoder(c.Request().Body).Decode(&sysCode)
	return c.JSON(http.StatusOK, codeService.UpdateByCode(sysCode, sysCodeDetails))
}

//DeleteSysCode :
func DeleteSysCode(c echo.Context) error {
	sysCode := c.Param("sys-code-id")
	return c.JSON(http.StatusOK, codeService.DeleteByCode(sysCode))
}
>>>>>>> d9b04e621bed3908f3ac63ebb93a04dae8c64749
