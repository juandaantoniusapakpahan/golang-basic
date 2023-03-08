package main

import "net/http"

var USERNAME = "coding"
var PASSWORD = "1234"


type CustomMux struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (c *CustomMux) RegisterMiddleware(next func(http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares{
		current = next(current)
	}

	current.ServeHTTP(w, r)
}


func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`Something wrong`))
			return
		}

		if USERNAME != username || PASSWORD != password {
			w.Write([]byte(`wrong password/username`))
			return
		}

		next.ServeHTTP(w, r)
	})
}


func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte(`Method should GET`))
			return
		}

		next.ServeHTTP(w, r)
	})
}