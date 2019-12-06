package main

import (
	"encoding/json"
	"gitlab.com/organisation/service"
	"net/http"
	"github.com/go-chi/chi"
	"log"
	"strconv"
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

func main()  {
	  log.Println("Server started on: http://localhost:8080")
	  var chiRouter = chi.NewRouter()
	  chiRouter.Get("/organizations", GetOrganisations)
	  chiRouter.Get("/organizations/{orgid}", GetOrganisationsById)
	  log.Fatal(http.ListenAndServe(":8080", chiRouter))
}
