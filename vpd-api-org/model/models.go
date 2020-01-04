package model

//Org : properties,
type Org struct {
 TenantID string
 Name string
 PrimaryContact string
 PrimaryContactPhone string
 PrimaryContactEmail string
 CreatedBy string
 Subscription BillingSubscription
 Plan BillingPlan
}

type BillingSubscription struct{
ID string
Name string
Description string

}

type BillingPlan struct{
ID string
Name string
Description string

}