package controller

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gitlab.com/vpd-payroll/vpd-api-sys/configuration"
	"gitlab.com/vpd-payroll/vpd-api-sys/repository"
	"gitlab.com/vpd-payroll/vpd-api-sys/services"
)

var db = configuration.InitializeDbCon()
var servRepository repository.RepositoryImpl
var servService services.ServiceImpl

//Sys Field Value
var valueRepository repository.SysValueRepositoryImpl
var valueService services.SysValueServiceImpl

//Sys Field Code
var codeRepository repository.SysCodeRepositoryImpl
var codeService services.SysCodeServiceImpl

// Country Field Code
var cntryRepository repository.CountryRepositoryImpl
var cntryService services.CountryServiceImpl

//Intilize :
func Intilize() {
	servRepository = repository.RepositoryImpl{Db: db}
	servService = services.ServiceImpl{Repo: servRepository}
	valueRepository = repository.SysValueRepositoryImpl{Db: db}
	valueService = services.SysValueServiceImpl{Repo: valueRepository}
	codeRepository = repository.SysCodeRepositoryImpl{Db: db}
	codeService = services.SysCodeServiceImpl{Repo: codeRepository}
	cntryRepository = repository.CountryRepositoryImpl{Db: db}
	cntryService = services.CountryServiceImpl{Repo: cntryRepository}
}

//Routing :
func Routing() {
	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.GET("/services", GetServ)
	e.GET("/features", GetFeatures)
	e.GET("/endpoints", Getendpoints)
	e.GET("/menu", GetMenu)
	e.POST("/menu/add", AddMenu)
	e.POST("/services", AddService)
	e.POST("/features", AddFeature)
	e.POST("/endpoints", AddEndpoint)
	e.PUT("/features/:id", UpdateFeature)
	//Sys Field Value
	e.POST("/sys-value", AddSysValue)
	//Sys Field Code
	e.POST("/sys-code", AddSysCode)

	//Country
	e.GET("/countries", GetCountries)
	e.GET("/states/:id", GetStates)
	e.GET("/cities/:id", GetCities)
	//Sys Field Value
	e.GET("/sys-field-values", GetAllSysValues)
	e.GET("/sys-field-value/:id", GetBySysValue)
	e.POST("/sys-field-value/add", AddSysValue)
	e.PUT("/sys-field-value/update/:id", UpdateSysValue)
	e.DELETE("/sys-field-value/delete/:id", DeleteSysValue)

	//Sys Field Code
	e.GET("/sys-field-codes", GetAllSysCodes)
	e.GET("/sys-field-code/:id", GetBySysCode)
	e.POST("/sys-field-code/add", AddSysCode)
	e.PUT("/sys-field-code/update/:id", UpdateSysCode)
	e.DELETE("/sys-field-code/delete/:id", DeleteSysCode)

	e.Logger.Fatal(e.Start(":8002"))
}
