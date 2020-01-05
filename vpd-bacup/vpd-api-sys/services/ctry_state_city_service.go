package services

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
	"gitlab.com/vpd-payroll/vpd-api-sys/repository"
)

type CountryService interface {
GetCountries()[]entities.Country
GetStates(ID string)[]entities.State
GetCities(ID string)[]entities.City
}

//CountryServiceImpl :
type CountryServiceImpl struct {
	Repo repository.CountryRepositoryImpl
}

//GetCountries
func(serv CountryServiceImpl)GetCountries()[]entities.Country{
return serv.Repo.GetCountries()
}

//GetStates
func(serv CountryServiceImpl)GetStates(ID string)[]entities.State{
 return serv.Repo.GetStates(ID)
}

//GetCities
func(serv CountryServiceImpl)GetCities(ID string)[]entities.City{
 return serv.Repo.GetCities(ID)
}

