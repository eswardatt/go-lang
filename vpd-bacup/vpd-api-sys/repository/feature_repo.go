package repository

import "gitlab.com/vpd-payroll/vpd-api-sys/entities"

func (repo RepositoryImpl) CreateFeature(f entities.Feature) {
	_, err := repo.Db.Query("Insert into sys.feature(feat_name, feat_description,serv_id) values($1,$2,$3)", f.Name, f.Description, f.Serv_id.ID)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	defer repo.Db.Close()
}

func (repo RepositoryImpl) GetFeatures() []entities.Feature {
	featureListT, err := repo.Db.Query("SELECT * FROM sys.feature")
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	feature := entities.Feature{}
	var features []entities.Feature
	for featureListT.Next() {
		errr := featureListT.Scan(&feature.ID, &feature.Name, &feature.Description, &feature.Serv_id.ID)
		if errr != nil {
			panic("Error while Reading  : " + errr.Error())
		}
		features = append(features, feature)
	}
	defer repo.Db.Close()
	return features
}

func (repo RepositoryImpl) UpdateFeature(f entities.Feature, ID string) {
	_, err := repo.Db.Query("update sys.feature set feat_name=$1,feat_description=$2 where feat_id=$3", f.Name, f.Description, ID)
	if err != nil {
		panic("Error while Executing Query  : " + err.Error())
	}
	defer repo.Db.Close()
}
