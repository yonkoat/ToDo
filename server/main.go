package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./DB/database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе:", err)
	}
	//отдаем все статичные файлы(а то есть html,css,js)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./front"))))
	//редирект с корня сайта
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		} else {
			http.NotFound(w, r)
		}

	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./front/index.html")
	})
	//Обработка страницы регистрации
	http.HandleFunc("/registrate", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./front/registrate.html")
	})
	http.HandleFunc("/registrating", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		NEW_USERNAME := r.FormValue("username")
		NEW_PASSWORD := r.FormValue("password")
		EMAIL := r.FormValue("email")

		//вставляем значиения в БД
		stmt, err := db.Prepare("INSERT INTO users('username','email','password') VALUES(?,?,?)")
		if err != nil {
			log.Println("ошибка подготовки запроса", err)
			http.Error(w, "ошибка сервера", http.StatusInternalServerError)
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(NEW_USERNAME, EMAIL, NEW_PASSWORD)

		if err != nil {
			log.Println("ОШИБКА ВСТАВКИ В БАЗУ ДАННЫХ", err)
			http.Error(w, "ОШИБКА ВСТАВКИ ПОЛЬЗОВАТЕЛЯ", http.StatusInternalServerError)
			return
		}

		fmt.Printf("REGISTRATION:\n USERNAME-%v\n Password-%v\n EMAIL-%v", NEW_USERNAME, NEW_PASSWORD, EMAIL)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})
	//обработка страницы логина для получения последующих post запросов
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/login" {
			http.ServeFile(w, r, "./front/login.html")
		} else {
			http.NotFound(w, r)
		}

	})
	http.HandleFunc("/login_in", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		Username := r.FormValue("username")
		Password := r.FormValue("password")

		fmt.Printf("LOGIN IN:\n USERNAME-%v\n Password-%v\n", Username, Password)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	})
	//запуск сервера
	http.ListenAndServe(":8080", nil)
}
