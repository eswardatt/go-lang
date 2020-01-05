package service

import (
	//"gitlab.com/vpd-payroll/vpd-api-iam/entities"
	"gitlab.com/vpd-payroll/vpd-api-iam/entities"
	"gitlab.com/vpd-payroll/vpd-api-iam/repository"
)

//UserService Interface
type RoleService interface {
	GetUsers() []entities.User                                  //GetUsers Method Declaration
	GetUserById(userID string) entities.User                    //GetUserById Method Declaration
	AddUser(userDetails entities.User) string                   //AddUser Method Declaration
	UpdateUser(userID string, userDetails entities.User) string //UpdateUser Method Declaration
	DeleteUser(userID string) string                            //DeleteUser Method	Declaration
}

//UserServiceImpl Implmentation
type UserServiceImpl struct {
	Repo repository.UserRepository
}

//GetUsers :
func (objUser UserServiceImpl) GetUsers() []entities.User {
	return objUser.Repo.Get()
}
func (objUser UserServiceImpl) GetUserById(userID string) entities.User {
	return objUser.Repo.GetById(userID)
}
func (objUser UserServiceImpl) AddUser(userDetails entities.User) string {

	return objUser.Repo.Post(userDetails)

}
func (objUser UserServiceImpl) UpdateUser(userID string, userDetails entities.User) string {
	return objUser.Repo.Put(userID, userDetails)
}
func (objUser UserServiceImpl) DeleteUser(userID string) string {
	return objUser.Repo.Delete(userID)
}
