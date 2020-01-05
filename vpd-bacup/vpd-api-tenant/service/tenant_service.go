package service

import (
	"gitlab.com/vpd-payroll/vpd-api-tenant/model"
	"gitlab.com/vpd-payroll/vpd-api-tenant/repository"
)

//TenantService interface to declare methods
type TenantService interface {
	GetTenants() []model.Tenant
	GetTenantsByID(TenantID string) model.Tenant
	SaveTenant(tenant model.Tenant) string
	UpdateTenants(tenant model.Tenant, ID string) string
	DeleteTenant(ID string) string
}

//TenantServiceImpl structure
type TenantServiceImpl struct {
	Repo repository.TenantRepository
}

//GetTenants returns all tenants
func (serv TenantServiceImpl) GetTenants() []model.Tenant {
	return serv.Repo.Get()
}

//GetTenantsByID return tenants by id
func (serv TenantServiceImpl) GetTenantsByID(TenantID string) model.Tenant {
	return serv.Repo.GetbyID(TenantID)
}

//SaveTenant insert the tenant data to db
func (serv TenantServiceImpl) SaveTenant(tenant model.Tenant) string {

	isExists := serv.Repo.GetByName(tenant.Name)
	if isExists == true {
		return tenant.Name + " " + "Already Exists"
	} else {
		return serv.Repo.Post(tenant)
	}
}


//UpdateTenants update tenant info
func (serv TenantServiceImpl) UpdateTenants(tenant model.Tenant, ID string) string {
	return serv.Repo.Put(tenant, ID)
}

//DeleteTenant deletes the tenant info
func (serv TenantServiceImpl) DeleteTenant(ID string) string {
	isExists := serv.Repo.GetTenantIDinSub(ID)
	if isExists == 0 {
		return serv.Repo.Delete(ID)
	} else {
		return "You can't delete this"
	}
}
