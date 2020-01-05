package service

import (
	"gitlab.com/vpd-payroll/vpd-api-tenant/model"
	"gitlab.com/vpd-payroll/vpd-api-tenant/repository"
)

//TenantServiceImplSub strucutre
type TenantServiceImplSub struct {
	Repo repository.TenantRepositorySub
}

//TenantServiceSub interface to declare methods
type TenantServiceSub interface {
	GetSubscriptions() []model.Subscription
	GetSubByTenant(ID string) []model.Subscription
	SubscribeService(sub model.Subscribe) string
	UnSubscribe(sub model.Subscribe)string
}

//GetSubscriptions returns all tenants
func (serv TenantServiceImplSub) GetSubscriptions() []model.Subscription {
	return serv.Repo.GetSub()
}

//GetSubByTenant returns single tenant
func (serv TenantServiceImplSub) GetSubByTenant(ID string) []model.Subscription {
	return serv.Repo.GetSubByID(ID)
}
//SubscribeService
func (serv TenantServiceImplSub)SubscribeService(sub model.Subscribe) string{
return serv.Repo.PostSub(sub)
}

//UnSubscribe unscribes the service
func (serv TenantServiceImplSub) UnSubscribe(sub model.Subscribe)string{
return serv.Repo.DeleteSub(sub)

}