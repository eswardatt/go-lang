package main

import (
	_ "github.com/lib/pq"
	"gitlab.com/vpd-payroll/vpd-api-sys/controller"
)

func main() {
	controller.Intilize()
	controller.Routing()
}
