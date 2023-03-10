package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct{
	Name string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "This is The main page")
	})
	r.Any("/user", func(ctx echo.Context) (err error) {
		u := new(User)

		if err = ctx.Bind(u); err != nil {
			return
		}

		return ctx.JSON(http.StatusOK, u)
	})

	fmt.Println("Server started at localhost:4000")
	r.Start(":4000")
}