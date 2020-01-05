package main

import "gitlab.com/vpd-payroll/vpd-api-iam/controller"

func main() {
	controller.Intialize()
	controller.Routers()
}
