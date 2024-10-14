package web_forum

import (
	"net/http"

	utils "web_forum/internal/utils"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err == nil {
		utils.S_essions.DeleteSession(cookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   1800,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
