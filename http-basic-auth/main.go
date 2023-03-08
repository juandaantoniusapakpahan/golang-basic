package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func OutputJSON(w http.ResponseWriter, o interface{}) {
	value, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(value)
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if !Auth(w, r) {return}
	if !AlloOnlyGet(w, r) {return}

	if id:=r.URL.Query().Get("id"); id!= "" {
		OutputJSON(w, SelectStudent(id))
		return
	}

	OutputJSON(w, GetStudents())
}

func main() {
	http.HandleFunc("/student", ActionStudent)

	server := new(http.Server)
	server.Addr = ":8080"

	fmt.Println("Server started at localhost:8080")
	server.ListenAndServe()
}