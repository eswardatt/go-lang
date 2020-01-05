package repository

import (
		"gitlab.com/vpd-payroll/vpd-api-tenant/modelTenantRepositorySub
tpe TenantRepositorySub interface {
DeleteSub(sub model.Subscribe) string
}
	 
	//Get returns all Sub scriptions
	func (repo TenantRepositoryImpl) GetSu() []model.Subscription {
	
	ows, err := repo.Db.Queryx("select tenant.tenant_id, tenant.tenant_Name,tenant.tenant_contact,COALESCE(subsc.sub_id,'NA') as subID, COALESCE(serv.serv_id,'NA')as serv_ID,COALESCE(serv.serv_name,'NA'),COALESCE(serv.serv_description,'NA') from tenant.tenant as tenant left join tenant.subsription as subsc on tenant.tenant_id=subsc.tenant_id left join sys.service as serv on subsc.service_id = serv.serv_id where tenant.tenant_status ='Active'")
tenant := model.Tenant{}
	service := model.Service{}
	subscription := model.Subscript ion{}  
res := []model.Subscription{}
	if err != nil {
	nic(err.Error())
	} 
 
or rows.Next() {
		err := rows.Scan(
	tenant.ID, &tennt.Name, &tenant.Contact,
		cription.ID,
	srvice.ID, &service.Name, &service.Description)
if err != nil {
			panic(err.Error))
		
			subscriptio n.TenantDetai ls = tenant
			subscription.ServceDetails = service
			res = append (res, subscrip tion)
		
			
		
  
//Get returns single Subscrip t ion
func  ( repo Tenant RepositoryImp) GetSubByID(ID string) []model.Subscription {
	s err := repo.Db.Queryx("select tenant.tenant_id, tenant.tenant_Name,tenant.tenant_contact,COALESCE(subsc.sub_id,'NA'), COALESCE(serv.serv_id,'NA'),COALESCE(serv.serv_name,'NA'),COALESCE(serv.serv_description,'NA') from tenant.tenant as tenant left join tenant.subsription as subsc on tenant.tenant_id=subsc.tenant_id left join sys.service as serv on subsc.service_id = serv.serv_id where tenant.tenant_id=$1", ID)
	tenant := odel.Tenant{}
	ervice := model.Service{}
res := []model.Subscription{}
	if err != nil {
		panic(err.Error())   
	} 
	
	for rows.N ext() {
		err := rows.Sc an(
	&tenant.ID, &tenant.Name, &tenant.Contact,
			&subscription.ID,
	service.ID, &sevice.Name, &service.Description)
		 != nil {
	aic(err.Error())
}
		subscription.TenntDetails = tenant
		iption.ServiceDetils = service
			res = appen d(res, subscr iption)
		}
	  
		res
			
		
//PostSub Subscribe service  
func (repo TenantRepositoryIm p l) Postub(sub model.Subscribe) string {
	_, e r r := repo.D b.Queryx("insrt into tenant.subsription(tenant_id,service_id)values($1,$2)",
	bTenantID, sub.ServiceID)
ar returnmsg string
	if err != il {
	panic(err.Error())
	returnmsg = "error"
		return returnmsg
	} 
	returnmsg = "Subscribed"
		return returnmsg
}

//DeleteSub Subscribion
func (repo TenantRepoitoryImpl) DeleteSub(sub model.Subscribe) string {
	_, err := repo.DbQueryx("update tenant.subsription set subscription_status='DeActive' where tenant_id=$1 and service_id=$2", sub.TenantID, sub.ServiceID)
	vr returnmsg string
	if err != nil {
		panic(err.Error))
	returnmsg = "error"
}
	returnmsg = "Unsubscribd successflly!"
	return returnmsg
}  
