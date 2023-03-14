package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type Student struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Grade int    `bson:"grade"`
}

func connect() (*mongo.Database, error) {
	clienOption := options.Client()
	clienOption.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clienOption)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return client.Database("database-std"), nil
}

func StudentFormHandlr(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/formStudent.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}	

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data Student

	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(data)
	return
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello World`))
	return
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/form-student", StudentFormHandlr)
	http.HandleFunc("/student", AddStudentHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
