package main

import (
	

	"gitlab.com/vpd-payroll/fw/microservice"
	
)

type ExampleConfig struct {
	ConfigVarInt    int    `envconfig:"CONFIG_VAR_INT"`
	ConfigVarString string `envconfig:"CONFIG_VAR_STRING"`
	ConfigVarBool   bool   `envconfig:"CONFIG_VAR_BOOL"`
}

var api = microservice.API{
	Name:        "Example",
	BaseURI:     "/api/example",
	ConfigVar:   new(ExampleConfig),
	BindRoutes:  routes,
	BeforeStart: beforeStart,
}


func beforeStart() error {
    DbInitialize()
	return nil
}

func main() {

	api.Start()

}
