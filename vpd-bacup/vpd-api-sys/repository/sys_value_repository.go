package repository

import (
	"github.com/jmoiron/sqlx"

	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

//SysValueRepository : Repository Interface
type SysValueRepository interface {
	Get() []entities.SysFieldValues
	GetByValue(sysValue string) entities.SysFieldValues
	Post(sysFieldDetails entities.SysFieldValues) string
	Put(sysValue string, SysValueDetails entities.SysFieldValues) string
	Delete(sysValue string) string
}

//SysValueRepositoryImpl : Repository Implementation structure
type SysValueRepositoryImpl struct {
	Db *sqlx.DB
}

//Get :
func (repoSysValue SysValueRepositoryImpl) Get() []entities.SysFieldValues {
	queryRes, err := repoSysValue.Db.Query(
		"SELECT FIELD_VALUE, FIELD_CODE, DESCRIPTION, FIELD_VALUE_TYPE, IN_USE FROM SYS.SYS_FIELD_VALUES")
	entities.CheckErrorCaused(err, "Executing Query")

	resArray := entities.SysFieldValues{}
	var resSlice []entities.SysFieldValues
	for queryRes.Next() {
		output := queryRes.Scan(&resArray.FieldValue, &resArray.FieldCode, &resArray.Description, &resArray.FieldValueType, &resArray.InUse)
		entities.CheckErrorCaused(output, "Reading columns")
		resSlice = append(resSlice, resArray)
	}
	return resSlice
}

//GetByValue :
func (repoSysValue SysValueRepositoryImpl) GetByValue(sysValue string) entities.SysFieldValues {
	queryRes, err := repoSysValue.Db.Query(
		"SELECT FIELD_VALUE, FIELD_CODE, DESCRIPTION, FIELD_VALUE_TYPE, IN_USE FROM SYS.SYS_FIELD_VALUES WHERE FIELD_VALUE = $1",
		sysValue)
	entities.CheckErrorCaused(err, "Executing Query")

	resArray := entities.SysFieldValues{}
	for queryRes.Next() {
		output := queryRes.Scan(&resArray.FieldValue, &resArray.FieldCode, &resArray.Description, &resArray.FieldValueType, &resArray.InUse)
		entities.CheckErrorCaused(output, "Reading columns")
	}
	return resArray
}

//Post : Method for Creating New Users
func (repoSysValue SysValueRepositoryImpl) Post(SysValueDetails entities.SysFieldValues) string {
	_, err := repoSysValue.Db.Query(
		"INSERT INTO SYS.SYS_FIELD_VALUES(FIELD_VALUE, FIELD_CODE, DESCRIPTION, FIELD_VALUE_TYPE, IN_USE) VALUES($1, $2, $3, $4, $5)",
		SysValueDetails.FieldValue, SysValueDetails.FieldCode, SysValueDetails.Description, SysValueDetails.FieldValueType, SysValueDetails.InUse)
	entities.CheckErrorCaused(err, "Executing Query")

	return SysValueDetails.FieldValue + "Inserted successfully"
}

//Put :
func (repoSysValue SysValueRepositoryImpl) Put(sysValue string, sysValueDetails entities.SysFieldValues) string {
	_, err := repoSysValue.Db.Query(
		"UPDATE SYS.SYS_FIELD_VALUE SET DESCRIPTION = $3, FIELD_VALUE_TYPE = $4, IN_USE = $5 WHERE FIELD_VALUE = $1 && FIELD_CODE = $2",
		sysValueDetails.FieldValue, sysValueDetails.FieldCode, sysValueDetails.Description, sysValueDetails.FieldValueType, sysValueDetails.InUse)

	entities.CheckErrorCaused(err, "Executing Query")

	return "Sys Field Value " + sysValueDetails.FieldValue + " Updated successfully"
}

//Delete :
func (repoSysValue SysValueRepositoryImpl) Delete(sysValue string) string {
	_, err := repoSysValue.Db.Query("UPDATE SYS.SYS_FIELD_VALUE SET STATUS = $2 WHERE FIELD_VALUE = $1", sysValue, 0) //0 -- DeActive  || 1 -- Active

	entities.CheckErrorCaused(err, "Executing Query")
	return "Sys Field Value " + sysValue + " Deleted successfully"
}
