package main

import "net/http"

func main() {
	//отдаем все статичные файлы(а то есть html,css,js)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./front"))))
	//редирект с корня сайта
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "./front/index.html", http.StatusSeeOther)
			return
		} else {
			http.NotFound(w, r)
		}

	})
	//обработка страницы логина для получения последующих post запросов
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			http.ServeFile(w, r, "./front/login.html")
		} else {
			http.NotFound(w, r)
		}

	})
	//запуск сервера
	http.ListenAndServe(":8080", nil)
}
