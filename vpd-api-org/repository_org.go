package main

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-api-org/model"
)

//OrgRepositoryImpl implement interface
type OrgRepositoryImpl struct {
	Db *sqlx.DB
}

type OrgRepository interface {
	Get() []model.Org
	GetById(Id string) model.Org
	Post(org model.Org) string
	Put(org model.Org, Id string) string
	Delete(Id string) string
}

//Get returns organisations list
func (repo OrgRepositoryImpl) Get() []model.Org {
	rows, err := repo.Db.Queryx("SELECT tenant_id,name,primary_contact from org.org")
	org := model.Org{}
	res := []model.Org{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&org.TenantID, &org.Name, &org.PrimaryContact)
		if err != nil {
			panic(err.Error())
		}
		
		res = append(res, org)
	}
	return res
}

//GetById returns org by id
func (repo OrgRepositoryImpl) GetById(Id string) model.Org {
	rows, err := repo.Db.Queryx("SELECT tenant_id,name,primary_contact from org.org where tenant_id=$1",Id)
	org := model.Org{}
	
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&org.TenantID, &org.Name, &org.PrimaryContact)
		if err != nil {
			panic(err.Error())
		}
		
	}
	return org
}

//Post Insert Org
func (repo OrgRepositoryImpl) Post(org model.Org) string {
	_, err := repo.Db.Queryx("insert into org.org(name,primary_contact,primary_contact_phone,primary_contact_email,subscription_id,plan_id,created_by,created_at)values($1,$2,$3,$4,$5,$6,$7,NOW())",org.Name,org.PrimaryContact,org.PrimaryContactPhone,org.PrimaryContactEmail,org.Subscription.ID,org.Plan.ID,org.CreatedBy, )
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = org.Name + " " + "Saved successflly!"
	return returnmsg
}


//Put Update Org
func (repo OrgRepositoryImpl) Put(org model.Org, Id string) string {
	_, err := repo.Db.Queryx("update org.org set name=$1,primary_contact=$2,primary_contact_phone=$3,primary_contact_email=$4,subscription_id=$5,plan_id=$6,updated_by=$7,updated_at=NOW() where tenant_id=$9", org.Name, org.PrimaryContact,org.PrimaryContactPhone,org.PrimaryContactEmail,org.Subscription.ID,org.Plan.ID,org.CreatedBy,Id)
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = org.Name + " " + "Updated successflly!"
	return returnmsg
}

//Delete Org
func (repo OrgRepositoryImpl) Delete(Id string) string {
	_, err := repo.Db.Queryx("update org.org set org_status='DeActive' where tenant_id=$1", Id)
	var returnmsg string
	if err != nil {
		panic(err.Error())
		returnmsg = "error"
		return returnmsg
	}
	returnmsg = "Deleted successflly!"
	return returnmsg
}