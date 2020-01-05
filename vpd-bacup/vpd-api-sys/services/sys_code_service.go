package services

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
	"gitlab.com/vpd-payroll/vpd-api-sys/repository"
)

//SysCodeService :
type SysCodeService interface {
	RetrieveAll() []entities.SysFieldCode
	RetrieveByCode(sysCode string) entities.SysFieldCode
	Insert(SysCodeDetails entities.SysFieldCode) string
	UpdateByCode(sysCode string, SysCodeDetails entities.SysFieldCode) string
	DeleteByCode(sysCode string) string
}

//SysCodeServiceImpl :
type SysCodeServiceImpl struct {
	Repo repository.SysCodeRepository
}

//RetrieveAll :
func (objSysCode SysCodeServiceImpl) RetrieveAll() []entities.SysFieldCode {
	return objSysCode.Repo.Get()
}

//RetrieveByCode :
func (objSysCode SysCodeServiceImpl) RetrieveByCode(sysCode string) entities.SysFieldCode {
	return objSysCode.Repo.GetByCode(sysCode)
}

//Insert :
func (objSysCode SysCodeServiceImpl) Insert(SysCodeDetails entities.SysFieldCode) string {
	isexists := objSysCode.Repo.GetByCode(SysCodeDetails.FieldCode)

	if isexists.FieldCode != SysCodeDetails.FieldCode {
		return objSysCode.Repo.Post(SysCodeDetails)
	} else {
		return "SysFieldCode alredy exists"
	}
}

//UpdateByCode :
func (objSysCode SysCodeServiceImpl) UpdateByCode(sysCode string, SysCodeDetails entities.SysFieldCode) string {
	return objSysCode.Repo.Put(sysCode, SysCodeDetails)
}

//DeleteByCode :
func (objSysCode SysCodeServiceImpl) DeleteByCode(sysCode string) string {
	return objSysCode.Repo.Delete(sysCode)
}
