package controller

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/vpd-payroll/vpd-api-tenant/model"
	"gitlab.com/vpd-payroll/vpd-api-tenant/repository"
	"gitlab.com/vpd-payroll/vpd-api-tenant/service"
	//"github.com/labstack/echo/middleware"
	"encoding/json"
	_ "github.com/lib/pq"
	"gitlab.com/vpd-payroll/vpd-api-tenant/configuration"
)

var db = configuration.DbConn()
var servRepository repository.TenantRepositoryImpl

var servService service.TenantServiceImpl
var servServiceSub service.TenantServiceImplSub

func Intilize() {
	servRepository = repository.TenantRepositoryImpl{Db: db}
	servService = service.TenantServiceImpl{Repo: servRepository}

	servServiceSub = service.TenantServiceImplSub{Repo: servRepository}
}
func Routers() {
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//routing
	e.GET("/tenants", GetTenants)
	e.GET("/tenants/:id", GetTenantsByID)
	e.POST("/tenants", SaveTenant)
	e.PUT("/tenants/:id", UpdateTenants)

	e.DELETE("/tenants/:id", DeleteTenant)

	//e.GET("/subscriptions", GetSubscriptions)
	e.GET("/subscriptions/:id", GetSubByTenant)
	e.POST("/subscriptions", SubscribeService)
	e.DELETE("/subscriptions", UnSubscribe)

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

func SaveTenant(c echo.Context) error {
	tenant := model.Tenant{}
	json.NewDecoder(c.Request().Body).Decode(&tenant)
	tenantSave := servService.SaveTenant(tenant)
	return c.JSON(http.StatusOK, tenantSave)
}

func UpdateTenants(c echo.Context) error {
	id := c.Param("id")
	tenant := model.Tenant{}
	json.NewDecoder(c.Request().Body).Decode(&tenant)
	updateTenant := servService.UpdateTenants(tenant, id)
	return c.JSON(http.StatusOK, updateTenant)
}

func DeleteTenant(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servService.DeleteTenant(id))

}

func GetSubscriptions(c echo.Context) error {
	return c.JSON(http.StatusOK, servServiceSub.GetSubscriptions())
}

func GetSubByTenant(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, servServiceSub.GetSubByTenant(id))
}

func SubscribeService(c echo.Context) error {
	sub := model.Subscribe{}
	json.NewDecoder(c.Request().Body).Decode(&sub)
	serviceSub := servServiceSub.SubscribeService(sub)
	return c.JSON(http.StatusOK, serviceSub)
}

func UnSubscribe(c echo.Context) error {

   sub:= model.Subscribe{}
   json.NewDecoder(c.Request().Body).Decode(&sub)
   serviceSub:= servServiceSub.UnSubscribe(sub)
   return c.JSON(http.StatusOK, serviceSub)
}







