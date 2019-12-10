package main


import (
	
	"gitlab.com/organisation/configuration"
	"gitlab.com/organisation/entities"
	"fmt"
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


//GetDepartments By Organisation and department
func GetDepartmentsByOrgIdandDepId(OrgId,DepId int)[]entities.Department {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT dep.DEP_ID,dep.Name,org.ORG_ID,org.Name FROM Department as dep join Organisation as org on dep.ORG_ID=org.ORG_ID  where dep.DEP_ID=? and dep.ORG_ID=?",  OrgId,DepId)
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





//AddDepartment
func AddDepartment(dep entities.Department,Id int)string{
     db := configuration.DbConn()
	_,  err := db.Query("insert into Department(Name,ORG_ID) values(?,?) ",dep.Name,Id)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}
	returnmsg   = dep.Name +   " " + "Saved successflly!"
	return returnmsg
}

//Update Department
func UpdateDepartment(dep entities.Department,OrgId,DepId int)string{
     db := configuration.DbConn()
	_,  err := db.Query("update Department set Name=? where ORG_ID=? and DEP_ID=?", dep.Name,OrgId,DepId)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}
	returnmsg   = dep.Name +   " " + "Updated successflly!"
	return returnmsg
}

//Delete Department
func DeleteDepartment(OrgId,DepId int) string{
db := configuration.DbConn()
    _, err:=db.Query("delete from Department where ORG_ID=? and DEP_ID=?",OrgId,DepId)
    var returnmsg string
    if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}

	returnmsg   = "Deleted successflly!"
	return returnmsg
}

//Get Employees
func GetEmployees(OrgId,DepId int) []entities.Employee {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT emp.EMP_ID,emp.FIRST_NAME,emp.LAST_NAME,emp.Age,EMAIL, dep.DEP_ID,dep.Name,org.ORG_ID,org.Name FROM Employee as emp join  Department as dep on emp.DEP_ID=dep.DEP_ID join Organisation as org on dep.ORG_ID=org.ORG_ID  where dep.DEP_ID=? and dep.ORG_ID=?", OrgId,DepId)
	if err != nil {
		panic(err.Error())
	}
	org := entities.Organisation{}
	emp:=entities.Employee{}
	dep:=entities.Department{}
	res := []entities.Employee{}
	for selDB.Next() {
		err = selDB.Scan(&emp.ID,&emp.FirstName,&emp.LastName,&emp.Age,&emp.Email,&dep.ID,&dep.Name,&org.ID,&org.Name)
		if err != nil {
			panic(err.Error())
		}
		 emp.DepartmentDetails=dep
		emp.DepartmentDetails.OrganisationDetails=org
	   
		res = append(res, emp)
	}
	
	return res
}


//Get Employees By Id
func GetEmployeesById(OrgId,DepId,EmpId int) []entities.Employee {
	db := configuration.DbConn()
	selDB, err := db.Query("SELECT emp.EMP_ID,emp.FIRST_NAME,emp.LAST_NAME,emp.Age,EMAIL, dep.DEP_ID,dep.Name,org.ORG_ID,org.Name FROM Employee as emp join  Department as dep on emp.DEP_ID=dep.DEP_ID join Organisation as org on dep.ORG_ID=org.ORG_ID  where dep.DEP_ID=? and dep.ORG_ID=? and emp.EMP_ID=?", OrgId,DepId,EmpId)
	if err != nil {
		panic(err.Error())
	}
	org := entities.Organisation{}
	emp:=entities.Employee{}
	dep:=entities.Department{}
	res := []entities.Employee{}
	for selDB.Next() {
		err = selDB.Scan(&emp.ID,&emp.FirstName,&emp.LastName,&emp.Age,&emp.Email,&dep.ID,&dep.Name,&org.ID,&org.Name)
		if err != nil {
			panic(err.Error())
		}
		 emp.DepartmentDetails=dep
		emp.DepartmentDetails.OrganisationDetails=org
	   
		res = append(res, emp)
	}
	
	return res
}


//AddEmployee
func AddEmployee(emp entities.Employee,DepId int)string{
     db := configuration.DbConn()
	_,  err := db.Query("insert into Employee(FIRST_NAME,LAST_NAME,Age,EMAIL,DEP_ID) values(?,?,?,?,?) ",emp.FirstName,emp.LastName,emp.Age,emp.Email,DepId)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}
	returnmsg   = "Saved successflly!"
	return returnmsg
}


//Update Employee
func UpdateEmployee(emp entities.Employee,EmpId int)string{
     db := configuration.DbConn()
	_,  err := db.Query("update employee set FIRST_NAME=?,LAST_NAME=?,Age=?,EMAIL=?,DEP_ID=? where EMP_ID=?",emp.FirstName,emp.LastName,emp.Age,emp.Email,emp.DepartmentDetails.ID,EmpId)
	var returnmsg string
	if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}
	returnmsg   =  "Updated successflly!"
	return returnmsg
}

//Delete Department
func DeleteEmployee(EmpId int) string{
db := configuration.DbConn()
    _, err:=db.Query("delete from employee where EMP_ID=?" ,EmpId)
    var returnmsg string
    if err != nil {
		returnmsg   = "entered into error"
		panic(err.Error())
	}

	returnmsg   = "Deleted successflly!"
	return returnmsg
}


func main(){
var emp entities.Employee
emp.FirstName="Raj"
emp.LastName="Kumar"
emp.Age=25
emp.Email="rajkumar@gmail.com"
emp.DepartmentDetails.ID=5
//fmt.Println(AddEmployee(emp,5))
//fmt.Println(UpdateDepartment(5,2))
//fmt.Println(UpdateEmployee(emp,3))
fmt.Println(DeleteEmployee(3))
}

