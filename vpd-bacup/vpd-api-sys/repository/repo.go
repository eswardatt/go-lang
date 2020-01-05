package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

type RepositoryImpl struct {
	Db *sqlx.DB
}

type Repository interface {
	// GetAll will return all Services
	GetAll() []entities.Service
	// CreateService will return a Service
	CreateService(entities.Service)
	// CreateFeature will return a Feature
	CreateFeature(entities.Feature)
	//GetFeatures will return all Features
	GetFeatures() []entities.Feature
	//GetFeatures will return all Features
	UpdateFeature(entities.Feature, string)
	//CreateEndpoint will return all EndPoint
	CreateEndpoint(entities.EndPoint)
	//GetEndpoints will return all EndPoints
	GetEndpoints() []entities.EndPoint
	//GetMenu will return menu list
	GetMenu() []entities.Menu
	//CreateMenu will return a menu
	CreateMenu(entities.Menu)
}
