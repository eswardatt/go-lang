package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/jmoiron/sqlx"
	_"github.com/lib/pq"
	"log"
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


func Initialize(){
    db:= DbConn()
	servRepository := OrgRepositoryImpl{Db: db}
	service = OrgServiceImpl{Repo:servRepository}
	log.Println(service)
}
	


//GetOrgList returns organisation data in json format
func GetOrgList(c echo.Context) error {
	return c.JSON(http.StatusOK, service.GetOrgList())
}


func routes(e *echo.Echo) (err error) {
	e.GET("/", GetOrgList)
	return
}