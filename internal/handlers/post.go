package web_forum

import (
	"net/http"

	utils "web_forum/internal/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
	session, Found := utils.S_essions.GetSession(cookie.Value)
	if !Found {
		http.Redirect(w, r, "/", http.StatusUnauthorized)
		return
	}
}
