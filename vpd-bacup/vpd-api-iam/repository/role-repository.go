package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-api-iam/entities"
)

//UserRepository Interface
type RoleRepository interface {
	Get() []entities.Role //GetUsers Method Declaration
	GetById(roleid string) entities.Role
	GetByName(email string) bool                         //GetUserById Method Declaration
	Post(roleDetails entities.Role) string               //AddUser Method Declaration
	Put(roleid string, userDetails entities.Role) string //UpdateUser Method Declaration
	Delete(roleid string) string                         //DeleteUser Method	Declaration
}

//UserRepositoryImpl : Repository Implementation
type RoleRepositoryImpl struct {
	Db *sqlx.DB
}

//GetUsers : Method for Retrieving All Users
func (repoRole RoleRepositoryImpl) Get() []entities.Role {
	queryRes, err := repoRole.Db.Query("select role_name,role_description from iam.role")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}

	roleArray := entities.Role{}
	var roleSlice []entities.Role
	for queryRes.Next() {
		output := queryRes.Scan(&roleArray.Name, &roleArray.Description)
		if output != nil {
			panic("Error while reading : " + output.Error())
		}
		roleSlice = append(roleSlice, roleArray)
	}
	return roleSlice
}

//GetUserByID : Method for Retrieve an User By ID
func (repoRole RoleRepositoryImpl) GetById(roleid string) entities.Role {
	queryRes, err := repoRole.Db.Query("select role_id,role_name,role_description from iam.role where role_id=$1", roleid)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	res := entities.Role{}

	for queryRes.Next() {
		data := queryRes.Scan(&res.ID, &res.Name, &res.Description)
		if data != nil {
			panic(data.Error())
		}

	}
	return res
}

//AddUser : Method for Creating New Roles
func (repoRole RoleRepositoryImpl) Post(role entities.Role) string {

	_, err := repoRole.Db.Query("insert into iam.role(role_name,role_description) values($1,$2)", role.Name, role.Description)
	if err != nil {
		panic(err.Error())
	}
	strmsg := role.Name + "inserted successfully"
	return strmsg
}

//UpdateUser : Method for Updating User by ID
func (repoRole RoleRepositoryImpl) Put(roleid string, role entities.Role) string {
	_, err := repoRole.Db.Query("update iam.role set role_name=$1,role_description=$2 where role_id=$3", role.Name, role.Description, roleid)
	if err != nil {
		panic(err.Error())
	}
	strmsg := role.Name + "updated successfully"
	return strmsg
}

//DeleteUser : Method for Deleting User by ID
func (repoRole RoleRepositoryImpl) Delete(roleid string) string {
	_, err := repoRole.Db.Query("update iam.role set status=$2 where role_id=$1", roleid, "DeActive")
	if err != nil {
		panic(err.Error())
	}
	defer repoRole.Db.Close()
	strmsg := " status updated successfully"
	return strmsg
}
func (repoRole RoleRepositoryImpl) GetByName(name string) bool {
	queryRes, err := repoRole.Db.Query("select count(role_id) from iam.role where role_name=$1", name)
	var rolecount int
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	for queryRes.Next() {
		err = queryRes.Scan(&rolecount)
	}
	if rolecount != 1 {
		return false
	} else {
		return true
	}

}
