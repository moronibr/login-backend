package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "../Frontend/index.html"

	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return

	}

	err = t.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

var userDB = map[string]string{
	"admin": "password",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login Successful")
	} else {
		fmt.Fprintf(w, "Login Failed")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {

	case "/login":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)

	default:
		fmt.Fprintf(w, "Hello World")
	}

}

func main() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8181", nil)
}
