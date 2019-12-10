package main

import (
	"encoding/json"
	"gitlab.com/organisation/service"
	"net/http"
	"github.com/go-chi/chi"
	"log"
	"strconv"
    "gitlab.com/organisation/entities"
)

func GetOrganisations(response http.ResponseWriter, request *http.Request) {
	 x:=service.GetOrganisations()
	 json.NewEncoder(response).Encode(x)
}

func GetOrganisationsById(response http.ResponseWriter, request *http.Request) {
	
	queryString := chi.URLParam(request, "orgid")
	id, _ := strconv.Atoi(queryString)

	x:=service.GetOrganisationsById(id)
	json.NewEncoder(response).Encode(x)
}    

func AddOrganisation(response http.ResponseWriter, request *http.Request){
        org:= entities.Organisation{}
	    json.NewDecoder(request.Body).Decode(&org)
	    x:= service.AddOrganisation(org)
		json.NewEncoder(response).Encode(x)
}


func UpdateOrganisation(response http.ResponseWriter, request *http.Request){
        org:= entities.Organisation{}
	    json.NewDecoder(request.Body).Decode(&org)
	    x:= service.UpdateOrganisation(org)
		json.NewEncoder(response).Encode(x)
}

func DeleteOrganisation(response http.ResponseWriter, request *http.Request) {
	queryString := chi.URLParam(request, "orgid")
	id, _ := strconv.Atoi(queryString)
	x:=service.DeleteOrganisation(id)
	json.NewEncoder(response).Encode(x)
}


func GetDepartmentsByOrgId(response http.ResponseWriter, request *http.Request) {
	
	queryString := chi.URLParam(request, "orgid")
	id, _ := strconv.Atoi(queryString)
	x:=service.GetDepartmentsByOrgId(id)
	json.NewEncoder(response).Encode(x)
}

func GetDepartmentsByOrgIdandDepId(response http.ResponseWriter, request *http.Request) {
	
	queryStringorgid := chi.URLParam(request, "orgid")
	queryStringdepid := chi.URLParam(request, "depid")
	orgid, _ := strconv.Atoi(queryStringorgid)
	depid, _ := strconv.Atoi(queryStringdepid)
	x:=service.GetDepartmentsByOrgIdandDepId(orgid,depid)
	json.NewEncoder(response).Encode(x)
}

func AddDepartment(response http.ResponseWriter, request *http.Request){
        queryStringorgid := chi.URLParam(request, "orgid")
		orgid, _ := strconv.Atoi(queryStringorgid)
        org:= entities.Department{}
	    json.NewDecoder(request.Body).Decode(&org)
	    x:= service.AddDepartment(org,orgid)
		json.NewEncoder(response).Encode(x)
}

func UpdateDepartment(response http.ResponseWriter, request *http.Request){
        queryStringorgid := chi.URLParam(request, "orgid")
	    queryStringdepid := chi.URLParam(request, "depid")
	    orgid, _ := strconv.Atoi(queryStringorgid)
	    depid, _ := strconv.Atoi(queryStringdepid)
        org:= entities.Department{}
	    json.NewDecoder(request.Body).Decode(&org)
	    x:= service.UpdateDepartment(org,orgid,depid)
		json.NewEncoder(response).Encode(x)
}


func DeleteDepartment(response http.ResponseWriter, request *http.Request) {
	   queryStringorgid := chi.URLParam(request, "orgid")
	   queryStringdepid := chi.URLParam(request, "depid")
	   orgid, _ := strconv.Atoi(queryStringorgid)
	   depid, _ := strconv.Atoi(queryStringdepid)
	   x:=service.DeleteDepartment(orgid,depid)
	   json.NewEncoder(response).Encode(x)
}

func main()  {
	  log.Println("Server started on: http://localhost:8080")
	  var chiRouter = chi.NewRouter()
	  chiRouter.Get("/organizations", GetOrganisations)
	  chiRouter.Get("/organization/{orgid}", GetOrganisationsById)
	  chiRouter.Post("/organization/post", AddOrganisation)
	  chiRouter.Put("/organization/put", UpdateOrganisation)
	  chiRouter.Delete("/organization/delete/{orgid}", DeleteOrganisation)
	  chiRouter.Get("/organization/departments/{orgid}", GetDepartmentsByOrgId)
	  chiRouter.Get("/organization/departments/{orgid}/{depid}", GetDepartmentsByOrgIdandDepId)
	  chiRouter.Post("/organization/{orgid}/departments", AddDepartment)
	  chiRouter.Put("/organization/{orgid}/departments/{depid}", UpdateDepartment)
	  chiRouter.Delete("/organization/{orgid}/departments/{depid}", DeleteDepartment)
	  log.Fatal(http.ListenAndServe(":8080", chiRouter))
}
