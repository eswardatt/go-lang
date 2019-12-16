package service

import (
	"gitlab.com/echo/vpd-payroll/vpd-api-tenant/repository"
	"gitlab.com/echo/vpd-payroll/vpd-api-tenant/model"
)

type Service interface {
	GetTenants() []model.Tenant
}

type ServiceImpl struct {
	Repo repository.Repository
}


//GetTenants returns all tenants
func (serv ServiceImpl)GetTenants()[]model.Tenant{
 return serv.Repo.Get()
}


//GetTenants ByID
func (serv ServiceImpl)GetTenantsByID(TenantID string)model.Tenant{
 return serv.Repo.GetbyID(TenantID)
}

//SaveTenants
func (serv ServiceImpl)SaveTenant(tenant model.Tenant)string{
 return serv.Repo.Post(tenant)
}

//UpdateTenants
func (serv ServiceImpl)UpdateTenants(tenant model.Tenant)string{
 return serv.Repo.Put(tenant)
}

//DeleteTenant
func (serv ServiceImpl)DeleteTenant(ID string)string{
 return serv.Repo.Delete(ID)
}









