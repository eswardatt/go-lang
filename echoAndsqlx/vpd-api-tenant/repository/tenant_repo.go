package repository
import (
"gitlab.com/echo/vpd-payroll/vpd-api-tenant/configuration"
"gitlab.com/echo/vpd-payroll/vpd-api-tenant/model"
"fmt"

)



//GetTenants returns all tenants
func GetTenants()[]model.Tenant{
    db := configuration.DbConn()
	rows, err := db.Queryx("SELECT * FROM tenant.tenant")
    tenant := model.Tenant{}
	res := []model.Tenant{}
    if err != nil {
        panic(err.Error())
    }
	for rows.Next() {
        err := rows.Scan(&tenant.ID,&tenant.Name,&tenant.Contact)
        if err != nil {
            panic(err.Error())
        }
		res=append(res,tenant)
    }
	
	return res
}

//GetTenantsByID returns single tenant
func GetTenantsByID(TenantID string)model.Tenant{
db := configuration.DbConn()
	rows, err := db.Queryx("SELECT * FROM tenant.tenant where tenant_id=$1",TenantID)
    tenant := model.Tenant{}
    if err != nil {
        panic(err.Error())
    }
	for rows.Next() {
        err := rows.Scan(&tenant.ID,&tenant.Name,&tenant.Contact)
        if err != nil {
            panic(err.Error())
        }
    }
	
	return tenant
}

func main(){

fmt.Println(GetTenantsByID("t2"))

}
