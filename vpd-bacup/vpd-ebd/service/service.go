package service

import (
   "gitlab.com/vpd-payroll/vpd-ebd/entities"
   "gitlab.com/vpd-payroll/vpd-ebd/repository"
)

type EmployeeServiceImpl struct {
	Repo repository.EmployeeRepository
}

type EmployeeService interface {
GetEmployeeBankDetails(ID string) []entities.EmployeeBankDetails
GetEmployeePFDetails(ID string) []entities.EmployeePFDetails
}

func (serv EmployeeServiceImpl) GetEmployeeBankDetails(ID string) []entities.EmployeeBankDetails {
	return serv.Repo.GetEmployeeBankDetails(ID)
}

func (serv EmployeeServiceImpl) GetEmployeePFDetails(ID string) []entities.EmployeePFDetails {
	return serv.Repo.GetEmployeePFDetails(ID)
}