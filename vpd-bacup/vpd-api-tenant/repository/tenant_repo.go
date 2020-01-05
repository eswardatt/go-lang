package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-api-tenant/model"
)

//TenantRepositoryImpl implement interface
type TenantRepositoryImpl struct {
	Db *sqlx.DB
}

//TenantRepository interface to declare methods
type TenantRepository interface {
	Get() []model.Tenant
	GetbyID(TenantID string) model.Tenant
	Post(tenant model.Tenant) string
	Put(tenant model.Tenant, ID string) string
	Delete(ID string) string
	GetTenantIDinSub(ID string) int
	GetByName(Name string) bool
}

//Get returns all tenants
func (repo TenantRepositoryImpl) Get() []model.Tenant {

	rows, err := repo.Db.Queryx("SELECT tenant_id,tenant_name,tenant_contact FROM tenant.tenant where tenant_status='Active'")
	tenant := model.Tenant{}
	res := []model.Tenant{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&tenant.ID, &tenant.Name, &tenant.Contact)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, tenant)
	}
	return res
}

//GetbyID returns single tenant
func (repo TenantRepositoryImpl) GetbyID(TenantID string) model.Tenant {
	rows, err := repo.Db.Queryx("SELECT tenant_id,tenant_name,tenant_contact FROM tenant.tenant where tenant_id=$1", TenantID)
	tenant := model.Tenant{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		data := rows.Scan(&tenant.ID, &tenant.Name, &tenant.Contact)
		if data != nil {
			panic(data.Error())
		}
	}

	return tenant
}

//Post insert the Tenant data
func (repo TenantRepositoryImpl) Post(tenant model.Tenant) string {
	_, err := repo.Db.Queryx("insert into tenant.tenant(tenant_name,tenant_contact)values($1,$2)", tenant.Name, tenant.Contact)
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = tenant.Name + " " + "Saved successflly!"
	return returnmsg
}

//Put update the Tenant data
func (repo TenantRepositoryImpl) Put(tenant model.Tenant, ID string) string {
	_, err := repo.Db.Queryx("update tenant.tenant set tenant_name=$1,tenant_contact=$2 where tenant_id=$3", tenant.Name, tenant.Contact, ID)
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = tenant.Name + " " + "Updated successflly!"
	return returnmsg
}

//Delete the data in db
func (repo TenantRepositoryImpl) Delete(ID string) string {
	_, err := repo.Db.Queryx("update tenant.tenant set tenant_status='DeActive' where tenant_id=$1", ID)
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = "Deleted successflly!"
	return returnmsg
}

//GetTenantIDinSub retrieving the tenant id
func (repo TenantRepositoryImpl) GetTenantIDinSub(ID string) int {

	selDB, err := repo.Db.Queryx("select count(tenant_id) from tenant.subsription where tenant_id=$1", ID)
	var TenantIDCount int
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		err = selDB.Scan(&TenantIDCount)
	}
	return TenantIDCount
}

//GetByName retrieves the name
func (repo TenantRepositoryImpl) GetByName(Name string) bool {

	selDB, err := repo.Db.Queryx("select tenant_name from tenant.tenant where tenant_name=$1", Name)
	if err != nil {
		panic(err.Error())
	}
	var tenantName string
	for selDB.Next() {
		err = selDB.Scan(&tenantName)
	}
	if tenantName == Name {
		return true
	} else {
		return false
	}

}
