package web_forum

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	utils "web_forum/internal/utils"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if password != confirmPassword {
		fmt.Fprintf(w, "Registration failed: Passwords do not match")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_, err = utils.Db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, string(hashedPassword))
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Registration failed: Username already exists")
		return
	}

	fmt.Fprintf(w, "Registration successful for user: %s", username)
}
