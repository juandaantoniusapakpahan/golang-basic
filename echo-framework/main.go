package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

func main() {
	e := echo.New()
	e.GET("/",func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	// .QueryParam()
	e.GET("/student", func(c echo.Context) error {
		name := c.QueryParam("name")
		data := fmt.Sprintf("%s",name)
		return c.String(http.StatusOK, data)
	});

	// .Param()
	e.GET("/student/:id", func(c echo.Context) error{
		id := c.Param("id")
		data := fmt.Sprintf("%s", id)
		return c.String(http.StatusOK, data)
	})


	// .HTML()
	e.GET("/html", func(c echo.Context) error {
		data := "<h1>Hello form</h1>"
		return c.HTML(http.StatusOK, data)
	})

	// .Redirect()
	e.GET("/redirect", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	// .JSON
	e.GET("/json", func(c echo.Context) error {
		data := M{"name":"Richard Joho", "age": 32, "address": "USA"}
		return c.JSON(http.StatusOK, data)
	})


	// .Param("*")
	e.GET("/param/:id/*", func(c echo.Context) error{
		data1 := c.Param("id")
		data2 := c.Param("*")

		data := fmt.Sprintf("Peserta No %s, selakan naik ke paggung %s", data1, data2)
		return c.String(http.StatusOK, data)
	})


	// .FormValue()
	e.POST("/formvalue", func(c echo.Context) error {
		name:= c.FormValue("name")
		message:= c.FormValue("message")
		data := fmt.Sprintf("%s: %s", name, strings.Replace(message, "/", "", 20))
		return c.String(http.StatusOK, data)
	})


	e.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	e.GET("/about", AboutHandler)
	e.GET("home", echo.WrapHandler(HomeHandler))

	e.Logger.Fatal(e.Start(":3000"))
}

var AboutHandler = echo.WrapHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`About page`));
}))

var ActionIndex = func(W http.ResponseWriter, r *http.Request) {
	W.Write([]byte("from action index"))
}

var HomeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action home"))
})