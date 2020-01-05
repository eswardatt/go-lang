package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gitlab.com/vpd-payroll/vpd-api-iam/entities"
)

//UserRepository Interface
type UserRepository interface {
	Get() []entities.User //GetUsers Method Declaration
	GetById(userID string) entities.User
	GetByEmail(email string) bool
	//GetUserById Method Declaration
	Post(userDetails entities.User) string               //AddUser Method Declaration
	Put(userID string, userDetails entities.User) string //UpdateUser Method Declaration
	Delete(userID string) string                         //DeleteUser Method	Declaration
}

//UserRepositoryImpl : Repository Implementation
type UserRepositoryImpl struct {
	Db *sqlx.DB
}

//GetUsers : Method for Retrieving All Users
func (repoUser UserRepositoryImpl) Get() []entities.User {
	queryRes, err := repoUser.Db.Query("select user_id,user_name,email,profilepic from iam.user	where status=$1 ", "Activate")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}

	userArray := entities.User{}
	var userSlice []entities.User
	for queryRes.Next() {
<<<<<<< HEAD
		output := queryRes.Scan(&userArray.Userdetails.ID,&userArray.Userdetails.Name, &userArray.Userdetails.Email, &userArray.Roledetails.Name)
=======
		output := queryRes.Scan(&userArray.ID, &userArray.Name, &userArray.Email, &userArray.Profilepic)
>>>>>>> d9b04e621bed3908f3ac63ebb93a04dae8c64749
		if output != nil {
			panic("Error while reading : " + output.Error())
		}
		userSlice = append(userSlice, userArray)
	}
	return userSlice
}

//GetUserByID : Method for Retrieve an User By ID
func (repoUser UserRepositoryImpl) GetById(userID string) entities.User {
	queryRes, err := repoUser.Db.Query("select user_id,user_name,email from iam.user where user_id =$1", userID)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	res := entities.User{}

	for queryRes.Next() {
		data := queryRes.Scan(&res.ID, &res.Name, &res.Email)
		if data != nil {
			panic(data.Error())
		}
	}
	return res
}

//AddUser : Method for Creating New Users
<<<<<<< HEAD
func (repoUser UserRepositoryImpl) Post(userrole entities.User_Role) string {
	_, err := repoUser.Db.Query("call iam.Adduser($1,$2,$3,$4)",userrole.Userdetails.Name, userrole.Userdetails.Email, userrole.Userdetails.Profilepic, userrole.Roledetails.Name)
=======
func (repoUser UserRepositoryImpl) Post(user entities.User) string {
	_, err := repoUser.Db.Query("insert into iam.user (user_name,email,profilepic) values($1,$2,$3)", user.Name, user.Email, user.Profilepic)
>>>>>>> d9b04e621bed3908f3ac63ebb93a04dae8c64749
	if err != nil {
		panic(err.Error())
	}
	strmsg := user.Name + "inserted successfully"
	return strmsg
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//UpdateUser : Method for Updating User by ID
func (repoUser UserRepositoryImpl) Put(userID string, user entities.User) string {

	_, err := repoUser.Db.Query("update iam.user set user_name=$1,email=$2,profilepic=$3 WHERE user_id=$4", user.Name, user.Email, user.Profilepic, userID)
	if err != nil {
		log.Fatal(err)
	}

	return " updated successfully"
}

//DeleteUser : Method for Deleting User by ID
func (repoUser UserRepositoryImpl) Delete(userID string) string {

	_, err := repoUser.Db.Query("update iam.user set status=$1 WHERE user_id=$2", "DeActivate", userID)
	if err != nil {
		log.Fatal(err)
	}
	return " Deactivated successfully"

}
func (repoUser UserRepositoryImpl) GetByEmail(email string) bool {
	queryRes, err := repoUser.Db.Query("select count(user_id) from iam.user where email=$1", email)
	var usercount int
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	for queryRes.Next() {
		err = queryRes.Scan(&usercount)
	}
	if usercount != 1 {
		return false
	} else {
		return true
	}

}
