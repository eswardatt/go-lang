package entities

//Organisation Properties
type Organisation struct {
	ID   int
	Name string
}

//Department Properties
type Department struct {
	ID, OrgID int
	Name      string
}

//Employee Properties
type Employee struct {
	ID, DeptID, Age            int
	FirstName, LastName, Email string
}
