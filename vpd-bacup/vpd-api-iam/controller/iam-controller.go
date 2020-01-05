package controller

import (
	"net/http"

	"github.com/labstack/echo"
<<<<<<< HEAD
	//"github.com/labstack/echo/middleware"
=======
	"github.com/labstack/echo/middleware"

	//s"github.com/labstack/echo/middleware"
>>>>>>> d9b04e621bed3908f3ac63ebb93a04dae8c64749
	_ "github.com/lib/pq"
	configuration "gitlab.com/vpd-payroll/vpd-api-iam/config"
	"gitlab.com/vpd-payroll/vpd-api-iam/entities"
	"gitlab.com/vpd-payroll/vpd-api-iam/repository"
	service "gitlab.com/vpd-payroll/vpd-api-iam/services"

	"encoding/json"
)

var db = configuration.DbConn()
var userRepository repository.UserRepositoryImpl

var userService service.UserServiceImpl
var roleRepository repository.RoleRepositoryImpl

var roleService service.RoleServiceImpl

func Intialize() {
	userRepository = repository.UserRepositoryImpl{Db: db}
	userService = service.UserServiceImpl{Repo: userRepository}
	roleRepository = repository.RoleRepositoryImpl{Db: db}
	roleService = service.RoleServiceImpl{Repo: roleRepository}
	//service.ServiceImpl{Repo: servRepository}
}
func Routers() {
	e := echo.New()
	e.Use(middleware.CORS())
	//e.Use(middleware.Logger())
	//e.Use(middleware.Recover())
	//routing
	//e.Use(middleware.CORS())
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUserById)
	e.POST("/users", AddUser)
	e.PUT("/users/:id", UpdateUser)

	e.DELETE("/users/:id", DeleteUser)
	e.GET("/roles", GetRoles)
	e.GET("/roles/:id", GetRoleById)
	e.POST("/roles", AddRole)
	e.PUT("/roles/:id", UpdateRole)

	e.DELETE("/roles/:id", DeleteRole)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, userService.GetUsers())
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, userService.GetUserById(id))
}

func AddUser(c echo.Context) error {
	user := entities.User{}
	json.NewDecoder(c.Request().Body).Decode(&user)
	userSave := userService.AddUser(user)
	return c.JSON(http.StatusOK, userSave)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := entities.User{}
	json.NewDecoder(c.Request().Body).Decode(&user)
	updateUser := userService.UpdateUser(id, user)
	return c.JSON(http.StatusOK, updateUser)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, userService.DeleteUser(id))

}

// role controler
func GetRoles(c echo.Context) error {
	return c.JSON(http.StatusOK, roleService.GetRoles())
}

func GetRoleById(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, roleService.GetRoleById(id))
}

func AddRole(c echo.Context) error {
	role := entities.Role{}
	json.NewDecoder(c.Request().Body).Decode(&role)
	roleSave := roleService.AddRole(role)
	return c.JSON(http.StatusOK, roleSave)
}

func UpdateRole(c echo.Context) error {
	id := c.Param("id")
	role := entities.Role{}
	json.NewDecoder(c.Request().Body).Decode(&role)
	updateRole := roleService.UpdateRole(role, id)
	return c.JSON(http.StatusOK, updateRole)
}

func DeleteRole(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, roleService.DeleteRole(id))

}
