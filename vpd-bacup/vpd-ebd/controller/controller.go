package controller


import (
    "net/http"
    "github.com/labstack/echo"
	_"github.com/lib/pq"
	"gitlab.com/vpd-payroll/vpd-ebd/configuration"
	"gitlab.com/vpd-payroll/vpd-ebd/repository"
	"gitlab.com/vpd-payroll/vpd-ebd/service"
	)
	
	
var db = configuration.DbConn()
var servRepository repository.EmployeeRepositoryImpl
var servService service.EmployeeServiceImpl

func Intilize() {
	servRepository = repository.EmployeeRepositoryImpl{Db: db}
	servService = service.EmployeeServiceImpl{Repo: servRepository}
}

func Routers() {
e := echo.New()

 //routing
e.GET("/emppfdetails/:id", GetEmployeePFDetails)
e.GET("/empbankdetails/:id", GetEmployeeBankDetails)
// Start server
	    e.Logger.Fatal(e.Start(":8000"))
}

func GetEmployeeBankDetails(c echo.Context) error {
    id := c.Param("id")
	return c.JSON(http.StatusOK, servService.GetEmployeeBankDetails(id))
}

func GetEmployeePFDetails(c echo.Context) error {
    id := c.Param("id")
	return c.JSON(http.StatusOK, servService.GetEmployeePFDetails(id))
}