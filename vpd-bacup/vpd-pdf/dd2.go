package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// e.GET("/", func(c echo.Context) error {
	// 	return c.File("../sample.pdf")
	// })
	e.GET("/", AppendValuesToPdf)
	e.Logger.Fatal(e.Start(":8002"))
}
func AppendValuesToPdf(c echo.Context) error {

	return c.File(WritePdf())

}
func WritePdf() string {

	return "../pdf/form.pdf"
}
