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

func AddOrganisation(org entities.Organisation)string  {
	return repository.AddOrganisation(org)
}






