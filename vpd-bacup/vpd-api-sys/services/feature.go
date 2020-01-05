package services

import "gitlab.com/vpd-payroll/vpd-api-sys/entities"

func (serv ServiceImpl) AddFeature(f entities.Feature) {
	serv.Repo.CreateFeature(f)
}

func (serv ServiceImpl) GetFeatures() []entities.Feature {
	return serv.Repo.GetFeatures()
}

func (serv ServiceImpl) UpdateFeature(f entities.Feature, ID string) {
	serv.Repo.UpdateFeature(f, ID)
}
