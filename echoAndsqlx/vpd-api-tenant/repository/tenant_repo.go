package repository
import (

"gitlab.com/echo/vpd-payroll/vpd-api-tenant/model"
	"github.com/jmoiron/sqlx"
	

)

type RepositoryImpl struct {
	Db *sqlx.DB
}

type Repository interface {
	Get()[]model.Tenant
	GetbyID(TenantID string)model.Tenant
	Post(tenant model.Tenant)string
	Put(tenant model.Tenant)string
	Delete(ID string)string
}



//Get returns all tenants
func (repo RepositoryImpl)Get()[]model.Tenant{
    
	rows, err := repo.Db.Queryx("SELECT tenant_id,tenant_name,tenant_contact FROM tenant.tenant where tenant_status='Active'")
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

//GetID returns single tenant
func (repo RepositoryImpl)GetbyID(TenantID string)model.Tenant{
	rows, err := repo.Db.Queryx("SELECT tenant_id,tenant_name,tenant_contact FROM tenant.tenant where tenant_id=$1",TenantID)
    tenant := model.Tenant{}
    if err != nil {
        panic(err.Error())
    }
	for rows.Next() {
        data := rows.Scan(&tenant.ID,&tenant.Name,&tenant.Contact)
        if data != nil {
            panic(data.Error())
        }
    }
	
	return tenant
}

//Post Tenant
func (repo RepositoryImpl)Post(tenant model.Tenant)string{
     _,err := repo.Db.Queryx("insert into tenant.tenant(tenant_name,tenant_contact)values($1,$2)",tenant.Name,tenant.Contact)
	 var returnmsg string
     if err != nil {
	     panic(err.Error())
        returnmsg="error"
		return returnmsg
     }
     returnmsg   = tenant.Name +   " " + "Saved successflly!"
     return returnmsg
}


//Put Tenant
func (repo RepositoryImpl)Put(tenant model.Tenant)string{
     _,err := repo.Db.Queryx("update tenant.tenant set tenant_name=$1,tenant_contact=$2 where tenant_id=$3",tenant.Name,tenant.Contact,tenant.ID)
	 var returnmsg string
     if err != nil {
	     panic(err.Error())
        returnmsg="error"
		return returnmsg
     }
     returnmsg   = tenant.Name +   " " + "Updated successflly!"
     return returnmsg
}

//Delete
func (repo RepositoryImpl)Delete(ID string)string{
_,err:=repo.Db.Queryx("update tenant.tenant set tenant_status='DeActive' where tenant_id=$1",ID)
var returnmsg string
     if err != nil {
	     panic(err.Error())
        returnmsg="error"
		return returnmsg
     }
     returnmsg   = "Deleted successflly!"
     return returnmsg
}