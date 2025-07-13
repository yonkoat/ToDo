package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, "index.html")

	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		Login := r.FormValue("name")
		Password := r.FormValue("password")

		fmt.Printf("login is kaka - %v and the pass is kaka too - %v", Login, Password)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "regis.html")
	})

	http.HandleFunc("/REGsubmit", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		NEW_LOGIN := r.FormValue("name")
		NEW_PASSWORD := r.FormValue("password")
		EMAIL := r.FormValue("email")

		fmt.Printf(`NEW: 
					login - %v  
					pass - %v
					email - %v`, NEW_LOGIN, NEW_PASSWORD, EMAIL)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	http.ListenAndServe(":8080", nil)
}
