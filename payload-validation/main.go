package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"  validate:"required"`
  	Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age"   validate:"gte=0,lte=100"`
}

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    return cv.validator.Struct(i)
}


func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.POST("/users", func(c echo.Context)  error {
		  u := new(User)
    if err := c.Bind(u); err != nil {
        return err
    }
    if err := c.Validate(u); err != nil {
        return err
    }

    return c.JSON(http.StatusOK, true)
	})

	fmt.Println("Server started at: localhost:8080")
    e.Logger.Fatal(e.Start(":8080"))
}