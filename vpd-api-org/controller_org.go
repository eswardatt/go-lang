package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
	"log"
	"gitlab.com/vpd-payroll/vpd-api-org/model"
	"encoding/json"
	"gitlab.com/vpd-payroll/fw/microservice"
)

var	service OrgServiceImpl
var servRepository OrgRepositoryImpl

func DbConn() (db *sqlx.DB) {
	db, err := sqlx.Connect("postgres", "user=postgres password=123 dbname=vpd_payroll_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}




func DbInitialize(){
    db:= DbConn()
	servRepository := OrgRepositoryImpl{Db: db}
	service = OrgServiceImpl{Repo:servRepository}
	//log.Println(servRepository)
}
	


//GetOrgList returns organisation data in json format
func GetOrgList(c echo.Context) (err error) {
	microservice.WriteResult(c, service.GetOrgList(), microservice.Completed, microservice.NoArgs)
	return
}

//GetOrgById
func GetOrgById(c echo.Context)(err error){
     id := c.Param("id")
	 microservice.WriteResult(c, service.GetOrgById(id), microservice.Completed, microservice.NoArgs)
	 return
}

//Post Org
func saveOrg(c echo.Context) error {
    UserRequest := model.Org{}
	json.NewDecoder(c.Request().Body).Decode(&UserRequest)
	orgSave := service.SaveOrg(UserRequest)
	return c.JSON(http.StatusOK, orgSave)
}
func GetSubscriptions(c echo.Context) (err error) {
	microservice.WriteResult(c, service.GetSubscriptions(), microservice.Completed, microservice.NoArgs)
	return
}
func GetPlans(c echo.Context) (err error) {
	microservice.WriteResult(c, service.GetPlans(), microservice.Completed, microservice.NoArgs)
	return
}

func routes(e *echo.Echo) (err error) {
	e.GET("/organisations", GetOrgList)
	e.GET("/organisations/:id", GetOrgById)
	e.GET("/billingsubscriptions", GetSubscriptions)
	e.GET("/billingplans", GetPlans)
	e.POST("/organisations/", saveOrg)
	return
}



