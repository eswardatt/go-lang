package repository

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

//SysCodeRepository : Repository Interface
type SysCodeRepository interface {
	Get() []entities.SysFieldCode
	GetByCode(sysCode string) entities.SysFieldCode
	Post(SysCodeDetails entities.SysFieldCode) string
	Put(sysCode string, SysCodeDetails entities.SysFieldCode) string
	Delete(sysCode string) string
}

//SysCodeRepositoryImpl : Repository Implementation structure
type SysCodeRepositoryImpl struct {
	Db *sqlx.DB
}

//Get :
func (repoSysCode SysCodeRepositoryImpl) Get() []entities.SysFieldCode {
	queryRes, err := repoSysCode.Db.Query("SELECT FIELD_CODE, DESCRIPTION, FIELD_CODE_TYPE, IN_USE FROM SYS.SYS_FIELD_CODE")
	entities.CheckErrorCaused(err, "Executing Query")

	resArray := entities.SysFieldCode{}
	var resSlice []entities.SysFieldCode
	for queryRes.Next() {
		output := queryRes.Scan(&resArray.FieldCode, &resArray.Description, &resArray.FieldCodeType, &resArray.InUse)
		entities.CheckErrorCaused(output, "Reading columns")
		resSlice = append(resSlice, resArray)
	}
	return resSlice
}

//GetByCode :
func (repoSysCode SysCodeRepositoryImpl) GetByCode(sysCode string) entities.SysFieldCode {
	queryRes, err := repoSysCode.Db.Query("SELECT FIELD_CODE, DESCRIPTION, FIELD_CODE_TYPE, IN_USE FROM SYS.SYS_FIELD_CODE WHERE FIELD_CODE = $1", sysCode)
	entities.CheckErrorCaused(err, "Executing Query")

	resArray := entities.SysFieldCode{}
	for queryRes.Next() {
		output := queryRes.Scan(&resArray.FieldCode, &resArray.Description, &resArray.FieldCodeType, &resArray.InUse)
		entities.CheckErrorCaused(output, "Reading columns")
	}
	return resArray
}

//Post : Method for Creating New syscode
func (repoSysCode SysCodeRepositoryImpl) Post(SysCodeDetails entities.SysFieldCode) string {
	_, err := repoSysCode.Db.Query("INSERT INTO SYS.SYS_FIELD_CODE(FIELD_CODE, DESCRIPTION, FIELD_CODE_TYPE, IN_USE) VALUES($1, $2, $3, $4)", SysCodeDetails.FieldCode, SysCodeDetails.Description, SysCodeDetails.FieldCodeType, SysCodeDetails.InUse)

	entities.CheckErrorCaused(err, "Executing Query")

	return "Sys Field Code " + SysCodeDetails.FieldCode + " Inserted successfully"
}

//Put :
func (repoSysCode SysCodeRepositoryImpl) Put(sysCode string, SysCodeDetails entities.SysFieldCode) string {
	_, err := repoSysCode.Db.Query(
		"UPDATE SYS.SYS_FIELD_CODE SET DESCRIPTION = $2, FIELD_CODE_TYPE = $3, IN_USE = $4 WHERE FIELD_CODE = $1",
		SysCodeDetails.FieldCode, SysCodeDetails.Description, SysCodeDetails.FieldCodeType, SysCodeDetails.InUse)

	entities.CheckErrorCaused(err, "Executing Query")

	return "Sys Field Code " + SysCodeDetails.FieldCode + " Updated successfully"
}

//Delete :
func (repoSysCode SysCodeRepositoryImpl) Delete(sysCode string) string {
	_, err := repoSysCode.Db.Query("UPDATE SYS.SYS_FIELD_CODE SET STATUS = $2 WHERE FIELD_CODE = $1", sysCode, 0) //0 -- DeActive  || 1 -- Active

	entities.CheckErrorCaused(err, "Executing Query")
	return "Sys Field Code " + sysCode + " Deleted successfully"
}
