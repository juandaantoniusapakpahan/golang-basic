package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("MiddlewareOne Sucess")
		return next(c)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("MiddlewareTwo Success")
		return next(c)
	}
}

func otheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Using http Handler")
		next.ServeHTTP(w, r)
	})
}

func main() {
	e := echo.New()

	// middleware::logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method:=${method}, uri=${uri}, status=${status}, remote_ip=${remote_ip}, time=${time_rfc3339_nano}\n",
	}))
	e.GET("/index", func(c echo.Context) error {
		fmt.Println("With Middleware")
		return c.JSON(http.StatusOK, true)
	})

	e.Use(middlewareOne)
	e.Use(middlewareTwo)
	e.Use(echo.WrapMiddleware(otheMiddleware)) // middleware non-schema echo meddleware
	e.Logger.Fatal(e.Start(":8080"))
}
