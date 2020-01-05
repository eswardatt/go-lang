package model

//Tenant structure to declare variables
type Tenant struct {
	ID, Name string
	Contact  int
}

//Subscription structure to declare variables
type Subscription struct {
	ID             string
	TenantDetails  Tenant
	ServiceDetails Service
}

//Service structure to declare variables
type Service struct {
	ID                string
	Name, Description string
}

//type Subscribe struct{
//TenantID,ServiceID string
//}


//Subscribe structure to declare variables
type Subscribe struct {
	TenantID, ServiceID, SubscribeID string
}

