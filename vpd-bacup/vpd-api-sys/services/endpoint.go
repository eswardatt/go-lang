package services

import "gitlab.com/vpd-payroll/vpd-api-sys/entities"

func (serv ServiceImpl) AddEndpoint(e entities.EndPoint) {
	serv.Repo.CreateEndpoint(e)
}

func (serv ServiceImpl) GetEndpoints() []entities.EndPoint {
	return serv.Repo.GetEndpoints()
}
