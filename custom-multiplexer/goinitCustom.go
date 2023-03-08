package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func OutputJSON(w http.ResponseWriter, o interface{}) {
	value, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(`Something wrong`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(value)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id:=r.URL.Query().Get("id"); id!=""{
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetAllStudent())
}


func main() {
	mux := new(CustomMux)

	mux.HandleFunc("/student", ActionStudent)
	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)
	

	var server = new(http.Server)
	server.Addr = ":8080"
	server.Handler = mux

	fmt.Println("Server started at Localhost: 8080")
	server.ListenAndServe()
}