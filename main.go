package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var Tmpl *template.Template

func init() {
	Tmpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	r := http.NewServeMux()
	r.HandleFunc("/", UserLogin)
	r.HandleFunc("/auth", UserAuth)

	defer http.ListenAndServe(":8080", r)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	Tmpl.ExecuteTemplate(w, "userLogin.html", nil)
}

func UserAuth(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	userName := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Printf("Username :%s\tPassword :%s\n", userName, password)

	if userName == "user" && password == "user" {
		Tmpl.ExecuteTemplate(w, "auth.html", userName)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
