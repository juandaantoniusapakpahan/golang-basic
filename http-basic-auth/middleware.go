package main

import "net/http"

const USERNAME = "coding"
const PASSWORD = "tanker"


func Auth(w http.ResponseWriter, r *http.Request) bool{
	username, password, ok := r.BasicAuth()

	if !ok {
		w.Write([]byte(`Something wrong`))
		return false
	}

	if USERNAME != username || PASSWORD != password {
		w.Write([]byte(`wrong username/password`))
		return false
	}

	return true
}

func AlloOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte(`Just GET METHOD`))
		return false
	}
	return true
}