package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	handler "web_forum/internal/handlers"
	utils "web_forum/internal/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var err error
	utils.Template, err = utils.Template.ParseGlob("../templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	utils.Db, err = sql.Open("sqlite3", "../internal/db/users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer utils.Db.Close()
	_, err = utils.Db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = utils.Db.Exec(`CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		owner TEXT UNIQUE NOT NULL,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		categories TEXT NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}
	utils.S_essions = utils.NewSessionManager()
	http.HandleFunc("/", handler.HandleDashboard)
	http.HandleFunc("/login", handler.HandleLogin)
	http.HandleFunc("/logout", handler.HandleLogout)
	http.HandleFunc("/static/", handler.ServeStitic)
	http.HandleFunc("/register", handler.HandleRegister)
	http.HandleFunc("/post", handler.CreatePost)
	http.HandleFunc("/account", handler.ServeHTML)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
