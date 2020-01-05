package services

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
	"gitlab.com/vpd-payroll/vpd-api-sys/repository"
)

//SysValueService :
type SysValueService interface {
	RetrieveAll() []entities.SysFieldValues
	RetrieveByValue(sysValue string) entities.SysFieldValues
	Insert(SysValueDetails entities.SysFieldValues) string
	UpdateByValue(sysValue string, SysValueDetails entities.SysFieldValues) string
	DeleteByValue(sysValue string) string
}

//SysValueServiceImpl :
type SysValueServiceImpl struct {
	Repo repository.SysValueRepository
}

//RetrieveAll :
func (objSysValue SysValueServiceImpl) RetrieveAll() []entities.SysFieldValues {
	return objSysValue.Repo.Get()
}

//RetrieveByCode :
func (objSysValue SysValueServiceImpl) RetrieveByValue(sysValue string) entities.SysFieldValues {
	return objSysValue.Repo.GetByValue(sysValue)
}

//AddSysFieldValue :
func (objSysValue SysValueServiceImpl) Insert(SysValueDetails entities.SysFieldValues) string {
	//isexists := objSysValue.Repo.GetByValue(userDetails.Userdetails.Email)
	isexists := false
	if isexists == false {
		return objSysValue.Repo.Post(SysValueDetails)
	} else {
		return "SysFieldValue alredy exists"
	}
}

//UpdateByValue :
func (objSysValue SysValueServiceImpl) UpdateByValue(sysValue string, sysValueDetails entities.SysFieldValues) string {
	return objSysValue.Repo.Put(sysValue, sysValueDetails)
}

//DeleteByValue :
func (objSysValue SysValueServiceImpl) DeleteByValue(sysValue string) string {
	return objSysValue.Repo.Delete(sysValue)
}
