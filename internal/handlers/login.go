package web_forum

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	utils "web_forum/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := utils.Db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			fmt.Fprintf(w, "Login failed: User not found")
		} else {
			log.Println(err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		fmt.Fprintf(w, "Login failed: Incorrect password")
		return
	}

	sessionID := utils.S_essions.CreateSession(username)
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   1800,
	})

	fmt.Fprintf(w, "Login successful for user: %s", username)
}
