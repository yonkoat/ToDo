package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "./" {
			http.Redirect(w, r, "/index.html", http.StatusSeeOther)
			return
		}
		http.FileServer(http.Dir("./")).ServeHTTP(w, r)

	})
	http.ListenAndServe(":8080", nil)

}
