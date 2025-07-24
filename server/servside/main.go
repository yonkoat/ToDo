package main

import "net/http"

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./front"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "./front/index.html", http.StatusSeeOther)
			return
		} else {
			http.NotFound(w, r)
		}

	})

	http.ListenAndServe(":8080", nil)
}
