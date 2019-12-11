package service

import (
	
	"gitlab.com/organisation/entities"
	"gitlab.com/organisation/repository"
	
	
	
)


//GetOrganisations service
func GetOrganisations() []entities.Organisation {
	return repository.GetOrganisations()
}
//GetOrganisationsById
func GetOrganisationsById(Id int)[]entities.Organisation{
	return repository.GetOrganisationsById(Id)
}

//AddOrganisation
func AddOrganisation(org entities.Organisation)string  {
   //var returnMsg string
   x:=repository.CheckNameExistorNot(org.Name)
   if x==true{
      returnMsg:=org.Name+" "+"Already Existed"
      return returnMsg
   }else{
     return repository.AddOrganisation(org)
   }

	
}
//UpdateOrganisation
func UpdateOrganisation(org entities.Organisation)string  {
	return repository.UpdateOrganisation(org)
}
//DeleteOrganisation
func DeleteOrganisation(Id int)string  {
	return repository.DeleteOrganisation(Id)
}


//GetDepartmentsByOrgId
func GetDepartmentsByOrgId(Id int)[]entities.Department{
  return repository.GetDepartmentsByOrgId(Id)
}

func GetDepartmentsByOrgIdandDepId(OrgId,DepId int)[]entities.Department{
return repository.GetDepartmentsByOrgIdandDepId(OrgId,DepId)
}


//AddDepartment
func AddDepartment(dep entities.Department,Id int)string{
return repository.AddDepartment(dep,Id)
}
//UpdateDepartment
func UpdateDepartment(dep entities.Department,OrgId,DepId int)string{
return repository.UpdateDepartment(dep,OrgId,DepId)
}


//DeleteDepartment
func DeleteDepartment(OrgId,DepId int)string  {
	return repository.DeleteDepartment(OrgId,DepId)
}
//GetEmployees
func GetEmployees(OrgId,DepId int)[]entities.Employee{
     return repository.GetEmployees(OrgId,DepId)
}

//GetEmployeesById

func GetEmployeesById(OrgId,DepId,EmpId int)[]entities.Employee{
return repository.GetEmployeesById(OrgId,DepId,EmpId)
}

//AddEmployee
func AddEmployee(emp entities.Employee,DepId int)string{
return repository.AddEmployee(emp,DepId)
}

//UpdateEmployee

func UpdateEmployee(emp entities.Employee,EmpId int)string{
return repository.UpdateEmployee(emp,EmpId)
}

//DeleteEmployee
func DeleteEmployee(EmpId int)string{
return repository.DeleteEmployee(EmpId)
}


//func main(){
//fmt.Println(GetEmployeesById(2,11,3))
//}



