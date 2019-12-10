package repository

import (
	
	"gitlab.com/organisation/configuration"
	"gitlab.com/organisation/entities"
	
	)

func GetOrganisations() []entities.Organisation {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT * FROM Organisation where Name !=''")
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



//GetOrganisationsById
func GetOrganisationsById(Id int)  []entities.Organisation {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT * FROM Organisation where ORG_ID=?",  Id)

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

//AddOrganisation
func AddOrganisation(org entities.Organisation) string{
	db := configuration.DbConn()
	_,  err := db.Query("insert into Organisation(Name) values(?) ", org.Name)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}

	returnmsg   = org.Name +   " " + "Saved successflly!"
	return returnmsg

}



//Update Organisation
func UpdateOrganisation(org entities.Organisation) string{
db := configuration.DbConn()
	_,  err := db.Query("update Organisation set Name=? where ORG_ID=?", org.Name,org.ID)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}

	returnmsg   = "Updated successflly!"
	return returnmsg
}


//Delete Organisation
func DeleteOrganisation(Id int) string{
db := configuration.DbConn()
    _, err:=db.Query("delete from Organisation where ORG_ID=?",Id)
    var returnmsg string
    if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}

	returnmsg   = "Deleted successflly!"
	return returnmsg
}

//CheckNameExistorNot
func CheckNameExistorNot(Name string )bool{
db := configuration.DbConn()
selDB, err:=db.Query("select Name from Organisation where Name=?",Name)
   if err != nil {
		panic(err.Error())
	}
	var orgName string
	for selDB.Next(){
	err = selDB.Scan(&orgName)
	}
	
	if orgName==Name{
	return true
	}else{
	return false
	}
	
}



//GetDepartments By Organisation
func GetDepartmentsByOrgId(Id int)[]entities.Department {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT dep.DEP_ID,dep.Name,org.ORG_ID,org.Name FROM Department as dep join Organisation as org on dep.ORG_ID=org.ORG_ID  where dep.ORG_ID=?",  Id)
	if err != nil {
		panic(err.Error())
	}
	dep := entities.Department{}
	org:= entities.Organisation{}
	res := []entities.Department{}
	
	for selDB.Next() {
		err = selDB.Scan(&dep.ID,&dep.Name,&org.ID,&org.Name)
		if err != nil {
			panic(err.Error())
		}
		dep.OrganisationDetails=org
		res = append(res, dep)
		
	}
	return res
}

//GetDepartments By Organisation and Department
func GetDepartmentsByOrgIdandDep(OrgId,DepId int)[]entities.Department {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT dep.DEP_ID,dep.Name,org.ORG_ID,org.Name FROM Department as dep join Organisation as org on dep.ORG_ID=org.ORG_ID  where dep.ORG_ID=? and dep.DEP_ID=?",  OrgId,DepId)
	if err != nil {
		panic(err.Error())
	}
	dep := entities.Department{}
	res := []entities.Department{}
	
	for selDB.Next() {
		err = selDB.Scan(&dep.ID,&dep.Name,&dep.OrganisationDetails.ID,&dep.OrganisationDetails.Name)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, dep)
		
	}
	return res
}




