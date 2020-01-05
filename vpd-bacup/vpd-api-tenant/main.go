package main
import "gitlab.com/vpd-payroll/vpd-api-tenant/controller"

func main() {
	controller.Intilize()
	controller.Routers()
}
