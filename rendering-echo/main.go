package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template // parsing and rendering template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug
	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {

	if t.debug {
		t.ReloadTemplates()
	}
	return t.template.ExecuteTemplate(w, name, data)

}

func main() {
	e := echo.New()
	e.Renderer = NewRenderer("./*.html", true)

	e.GET("/index", func(c echo.Context) error {
		data := M{"message": "Welcome to Paradise"}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
