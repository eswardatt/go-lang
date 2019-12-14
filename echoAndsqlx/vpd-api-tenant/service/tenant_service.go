package service

import (
	"gitlab.com/echo/vpd-payroll/vpd-api-tenant/repository"
	"gitlab.com/echo/vpd-payroll/vpd-api-tenant/model"
	"fmt"
)


//GetTenants returns all tenants
func GetTenants()[]model.Tenant{
 return repository.GetTenants()
}


//GetTenantsByID returns single tenant
func GetTenantsByID(TenantID string)model.Tenant{
 return repository.GetTenantsByID(TenantID)
}

func main(){
fmt.Println(GetTenants())
fmt.Println(GetTenantsByID("t2"))
}

