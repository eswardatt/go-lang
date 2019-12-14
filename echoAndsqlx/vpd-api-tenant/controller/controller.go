package main

import (
"gitlab.com/echo/vpd-payroll/vpd-api-tenant/service"

"net/http"
"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"


)

func GetTenants(c echo.Context) error {
	return c.JSON(http.StatusOK, service.GetTenants())
}

func main(){
	e := echo.New()
        // Middleware
	    e.Use(middleware.Logger())
	    e.Use(middleware.Recover())
		// Routes
	    e.GET("/tenants/", GetTenants)
	
		// Start server
	    e.Logger.Fatal(e.Start(":8000"))

}
