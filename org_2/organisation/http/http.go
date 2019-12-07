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
	service.AddOrganisation(org)
}


func main()  {
	  log.Println("Server started on: http://localhost:8080")
	  var chiRouter = chi.NewRouter()
	  chiRouter.Get("/organizations", GetOrganisations)
	  chiRouter.Get("/organizations/{orgid}", GetOrganisationsById)
	  //chiRouter.Post("/organizations/", AddOrganisation)
	  log.Fatal(http.ListenAndServe(":8080", chiRouter))
}
