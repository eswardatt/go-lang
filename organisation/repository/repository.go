package repository

import (
	
	"gitlab.com/organisation/configuration"
	"gitlab.com/organisation/entities"


)

func GetOrganisations() []entities.Organisation {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT * FROM Organisation")
	if err != nil {
		panic(err.Error())
	}
	org := entities.Organisation{}
	res := []entities.Organisation{}
	for selDB.Next() {
		err = selDB.Scan(&org.ID, &org.Name)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, org)
	}
	return res
}

func GetOrganisationsById(Id int)[]entities.Organisation {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT * FROM Organisation where ORG_ID=?",Id)

	if err != nil {
		panic(err.Error())
	}
	org := entities.Organisation{}
	res := []entities.Organisation{}
	for selDB.Next() {
		err = selDB.Scan(&org.ID, &org.Name)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, org)
	}
	return res
}


