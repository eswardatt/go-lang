package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	// Read form fields
	//name := c.FormValue("name")
	//email := c.FormValue("email")

	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	//abs, _ := filepath.Abs("../filefolder")

	//dst, err := os.Create(abs + "/" + file.Filename)

	//dst, err := os.Create(file.Filename)
	dst, err := os.Create(filepath.Join("filefolder", filepath.Base(file.Filename))) // dir is directory where you want to save file.

	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully </p>", file.Filename))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.Static("/filefolder", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}
