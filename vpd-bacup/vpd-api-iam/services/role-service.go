package service

import (
	//"gitlab.com/vpd-payroll/vpd-api-iam/entities"
	"gitlab.com/vpd-payroll/vpd-api-iam/entities"
	"gitlab.com/vpd-payroll/vpd-api-iam/repository"
)

//UserService Interface
type UserService interface {
	GetRoles() []entities.Role                //GetUsers Method Declaration
	GetRoleById(roleID string) entities.Role  //GetUserById Method Declaration
	AddRole(roleDetails entities.Role) string //AddUser Method Declaration
	UpdateRole(roleID string) string          //UpdateUser Method Declaration
	DeleteRole(roleID string) string          //DeleteUser Method	Declaration
}

//UserServiceImpl Implmentation
type RoleServiceImpl struct {
	Repo repository.RoleRepository
}

//GetUsers :
func (objRole RoleServiceImpl) GetRoles() []entities.Role {
	return objRole.Repo.Get()
}
func (objRole RoleServiceImpl) GetRoleById(roleID string) entities.Role {
	return objRole.Repo.GetById(roleID)
}
func (objRole RoleServiceImpl) AddRole(role entities.Role) string {
	isexists := objRole.Repo.GetByName(role.Name)
	if isexists == false {
		return objRole.Repo.Post(role)
	} else {
		return "user alredy exists"
	}
}
func (objRole RoleServiceImpl) UpdateRole(role entities.Role, roleID string) string {
	return objRole.Repo.Put(roleID, role)
}
func (objRole RoleServiceImpl) DeleteRole(roleID string) string {
	return objRole.Repo.Delete(roleID)
}
