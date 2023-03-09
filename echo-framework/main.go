package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/student", func(c echo.Context) error {
		name := c.QueryParam("name")
		data := fmt.Sprintf("%s",name)
		return c.String(http.StatusOK, data)
	});
	e.GET("/student/:id", func(c echo.Context) error{
		id := c.Param("id")
		data := fmt.Sprintf("%s", id)
		return c.String(http.StatusOK, data)
	})

	e.GET("/html", func(c echo.Context) error {
		data := "<h1>Hello form</h1>"
		return c.HTML(http.StatusOK, data)
	})
	e.Logger.Fatal(e.Start(":3000"))
}