package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		http.FileServer(http.Dir("./")).ServeHTTP(w, r)

	})
	http.ListenAndServe(":8080", nil)

}
