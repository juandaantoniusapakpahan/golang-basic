package main

import (
	"fmt"
	"net/http"
	"time"

	gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var cookieName = "CookieData"

func DeleteHandler(w http.ResponseWriter, r *http.Request){
	c :=&http.Cookie{}
	c.Name = cookieName
	c.Expires= time.Unix(0,0)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cookieName := "CookieData"

	c := &http.Cookie{}

	if storedCookie, _ := r.Cookie(cookieName); storedCookie != nil {
		c = storedCookie
	}

	if c.Value == "" {
		c = &http.Cookie{}
		c.Name = cookieName
		c.Value = gubrak.RandomString(32)
		c.Expires =time.Now().Add(5 * time.Minute)
		http.SetCookie(w, c)
	}
	w.Write([]byte(c.Value))

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/delete", DeleteHandler)

	var server = new(http.Server)
	server.Addr =":8080"
	
	fmt.Println("Server started at localhost:8080")
	server.ListenAndServe()
}