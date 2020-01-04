package main

import "gitlab.com/vpd-payroll/vpd-api-org/model"

type OrgServiceImpl struct {
	Repo OrgRepository
}

type OrgService interface {
	GetOrgList() []model.Org
	GetOrgById(Id string) model.Org
	SaveOrg(org model.Org) string
	UpdateOrg(org model.Org, Id string) string
	DeleteOrg(Id string) string
}

//GetTenants returns all tenants
func (serv OrgServiceImpl) GetOrgList() []model.Org {
	return serv.Repo.Get()
}

//GetTenantsByID return tenants by id
func (serv OrgServiceImpl) GetOrgById(Id string) model.Org {
	return serv.Repo.GetById(Id)
}

//Post Organisation
func (serv OrgServiceImpl) SaveOrg(org model.Org) string {
	return serv.Repo.Post(org)
}

func (serv OrgServiceImpl) UpdateOrg(org model.Org, Id string) string {
	return serv.Repo.Put(org, Id)
}

func (serv OrgServiceImpl) DeleteOrg(Id string) string {
	return serv.Repo.Delete(Id)
}
