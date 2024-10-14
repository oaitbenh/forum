package web_forum

import (
	"fmt"
	"net/http"

	utils "web_forum/internal/utils"
)

func HandleDashboard(w http.ResponseWriter, r *http.Request) {
	var session utils.Session
	cookie, err := r.Cookie("session_id")
	if err == nil {
		session, _ = utils.S_essions.GetSession(cookie.Value)
	}
	fmt.Println(session)
	user := User{Username: session.Username}
	if session.Username != "" {
		user.Found = true
	}
	err = utils.Template.ExecuteTemplate(w, "HomePage.html", user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
