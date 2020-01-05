package repository

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

//CountryRepositoryImpl implement interface
type CountryRepositoryImpl struct {
	Db *sqlx.DB
}

type CountryRepository interface {
	GetCountries() []entities.Country
	GetStates(ID string) []entities.State
	GetCities(ID string) []entities.City
}

func (repo CountryRepositoryImpl) GetCountries() []entities.Country {
	rows, err := repo.Db.Queryx("SELECT id,shortname,name,phonecode FROM sys.sys_country")
	country := entities.Country{}
	res := []entities.Country{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&country.ID, &country.ShortName, &country.Name, &country.PhoneCode)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, country)
	}
	return res
}

func (repo CountryRepositoryImpl) GetStates(ID string) []entities.State {
	rows, err := repo.Db.Queryx("SELECT id,name,country_id FROM sys.sys_state where country_id=$1", ID)
	state := entities.State{}
	res := []entities.State{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&state.ID, &state.Name, &state.CountryID)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, state)
	}
	return res
}

func (repo CountryRepositoryImpl) GetCities(ID string) []entities.City {
	rows, err := repo.Db.Queryx("SELECT id,name,state_id FROM sys.sys_cities where state_id=$1", ID)
	city := entities.City{}
	res := []entities.City{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&city.ID, &city.Name, &city.StateID)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, city)
	}
	return res
}
