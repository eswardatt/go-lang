package services

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func (serv ServiceImpl) Get() []entities.Service {
	return serv.Repo.GetAll()
}
func (serv ServiceImpl) AddService(objserv entities.Service) {
	serv.Repo.CreateService(objserv)
}
