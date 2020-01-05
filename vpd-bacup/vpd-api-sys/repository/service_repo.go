package repository

import (
	"gitlab.com/vpd-payroll/vpd-api-sys/entities"
)

func (repo RepositoryImpl) GetAll() []entities.Service {
	serviceList, err := repo.Db.Query("SELECT serv_id, serv_name, serv_description FROM sys.service")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	service := entities.Service{}
	var services []entities.Service
	for serviceList.Next() {
		errr := serviceList.Scan(&service.ID, &service.Name, &service.Description)
		if errr != nil {
			panic("Error while Reading  : " + errr.Error())
		}
		services = append(services, service)
	}
	defer repo.Db.Close()

	return services
}

func (repo RepositoryImpl) CreateService(objserv entities.Service) {
	_, err := repo.Db.Query("Insert into sys.service(serv_name, serv_description) values($1,$2)", objserv.Name, objserv.Description)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	defer repo.Db.Close()

}
