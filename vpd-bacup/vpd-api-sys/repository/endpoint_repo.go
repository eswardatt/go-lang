package repository

import "gitlab.com/vpd-payroll/vpd-api-sys/entities"

func (repo RepositoryImpl) CreateEndpoint(e entities.EndPoint) {
	_, err := repo.Db.Query("Insert into sys.endpoint(endp_name, endp_description,feat_id) values($1,$2,$3)", e.Name, e.Description, e.Feat_id.ID)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	repo.Db.Close()

}

func (repo RepositoryImpl) GetEndpoints() []entities.EndPoint {
	endpointListT, err := repo.Db.Query("SELECT * FROM sys.endpoint")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	endpoint := entities.EndPoint{}
	var endpoints []entities.EndPoint
	for endpointListT.Next() {
		errr := endpointListT.Scan(&endpoint.ID, &endpoint.Name, &endpoint.Description, &endpoint.Feat_id.ID)
		if errr != nil {
			panic("Error while Reading  : " + errr.Error())
		}
		endpoints = append(endpoints, endpoint)
	}
	repo.Db.Close()

	return endpoints
}
