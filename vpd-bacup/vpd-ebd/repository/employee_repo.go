
package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-ebd/entities"
)

type EmployeeRepositoryImpl struct {
	Db *sqlx.DB
}

type EmployeeRepository interface {
GetEmployeeBankDetails(ID string) []entities.EmployeeBankDetails
GetEmployeePFDetails(ID string) []entities.EmployeePFDetails
}

func (repo EmployeeRepositoryImpl) GetEmployeeBankDetails(ID string) []entities.EmployeeBankDetails {
	rows, err := repo.Db.Queryx("SELECT * from employee.emp_bank_details where employee_number=$1",ID)
	emp := entities.EmployeeBankDetails{}
	res := []entities.EmployeeBankDetails{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&emp.EmployeeNumber,&emp.BankAccountNumber,&emp.IfscCode,&emp.BankAccountType,
		&emp.BankName,
		&emp.BranchName,&emp.SalaryPyamentMode,&emp.NameAsPerBank,&emp.Iban,&emp.Dd_payable_at)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, emp)
	}
	return res
}

func (repo EmployeeRepositoryImpl) GetEmployeePFDetails(ID string) []entities.EmployeePFDetails {

	rows, err := repo.Db.Queryx("SELECT * from employee.emp_pf_esidetails as epf join tenant.tenant as ten on ten.tenant_id=epf.tenant_id  where employee_id=$1",ID)
	emp := entities.EmployeePFDetails{}
	res := []entities.EmployeePFDetails{}
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		err := rows.Scan(&emp.TenantID,&emp.EmployeeID,&emp.IsEmployeeEligibleforPf,&emp.UniversalAccountNumber,
		&emp.PFNumber,
		&emp.PFScheme,&emp.PFJoiningDate,&emp.FamilyPfNumber,&emp.IsEpfContribution,&emp.IsEligibleEsi,&emp.EsiNumber,
		&emp.IsExistingMemberOfEps,&emp.TenanatName,&emp.TenanatContact)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, emp)
	}
	return res
}
