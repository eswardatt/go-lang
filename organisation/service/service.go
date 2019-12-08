package main

import (
	
	"gitlab.com/organisation/entities"
	"gitlab.com/organisation/repository"
	"fmt"
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
return GetDepartmentsByOrgId(Id)
}

func main(){
fmt.Println(GetDepartmentsByOrgId(2))
}



