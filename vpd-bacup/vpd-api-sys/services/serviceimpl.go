package services

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
	"gitlab.com/vpd-payroll/vpd-api-sys/repository"
)

type Service interface {
	//Get returns all Services
	Get() []entities.Service
	AddService(entities.Service)
	GetFeatures() entities.Feature
	AddFeature(entities.Feature)
	UpdateFeature(entities.Feature, string)
	GetEndpoints() []entities.EndPoint
	AddEndpoint(entities.EndPoint)

	GetMenu() []entities.Menu
	AddMenu(entities.Menu)
}

type ServiceImpl struct {
	Repo repository.Repository
}
