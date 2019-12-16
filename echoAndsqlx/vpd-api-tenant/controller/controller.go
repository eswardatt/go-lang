package main

import (
    "gitlab.com/echo/vpd-payroll/vpd-api-tenant/service"
	"gitlab.com/echo/vpd-payroll/vpd-api-tenant/repository"
	
	
    "net/http"
    "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
		_ "github.com/lib/pq"
			"gitlab.com/echo/vpd-payroll/vpd-api-tenant/configuration"



)
var db = configuration.DbConn()
var servRepository repository.RepositoryImpl

var servService service.ServiceImpl

func intilize() {
	servRepository = repository.RepositoryImpl{Db: db}
	servService = service.ServiceImpl{Repo: servRepository}
}
func routers() {
	e := echo.New()
	 e.Use(middleware.Logger())
	 e.Use(middleware.Recover())
	 //routing
	e.GET("/tenants", GetTenants)
	e.GET("/tenants/:id", GetTenantsByID)
	e.DELETE("/tenants/:id", DeleteTenant)

	// Start server
	    e.Logger.Fatal(e.Start(":8000"))
}


func GetTenants(c echo.Context) error {
	return c.JSON(http.StatusOK, servService.GetTenants())
}

func GetTenantsByID(c echo.Context) error {
    id := c.Param("id")
	return c.JSON(http.StatusOK, servService.GetTenantsByID(id))
}





func DeleteTenant(c echo.Context) error {
    id := c.Param("id")
	return c.JSON(http.StatusOK, servService.DeleteTenant(id))
}




func main(){
	intilize()
	routers()


}
