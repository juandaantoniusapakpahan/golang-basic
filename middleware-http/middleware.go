package main

import "net/http"

const USERNAME = "coding"
const PASSWORD = "tanker"


func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, err := r.BasicAuth()

		if !err {
			w.Write([]byte(`Something broke`))
			return
		}

		if USERNAME != username || PASSWORD != password {
			w.Write([]byte(`Wrong password/username`))
		}

		next.ServeHTTP(w, r)
	})
}


func MiddlewareAllowOnlyGet(next http.Handler)  http.Handler{
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte(`Just GET METHOD`))
		return
	}
	next.ServeHTTP(w, r)
})
}