package entities

//User structure
type User struct {
	ID             string
	Name           string
	Email          string
	Profilepic     string
	IsActiveStatus string
}
type Role_Feature struct {
	RoleDetails Role
	FeatureId   string
}

//Role structure
type Role struct {
	ID             string
	Name           string
	Description    string
	IsActiveStatus string
}

type User_Role struct {
	Roledetails Role
	Userdetails User
}
