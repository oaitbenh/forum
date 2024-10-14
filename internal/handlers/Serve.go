package web_forum

import (
	"net/http"
	"os"
	"strings"

	utils "web_forum/internal/utils"
)

func ServeHTML(w http.ResponseWriter, r *http.Request) {
	err := utils.Template.ExecuteTemplate(w, "LoginRegister.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ServeStitic(w http.ResponseWriter, r *http.Request) {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("../static/")))
	_, err := os.Stat("../" + r.URL.Path)
	if strings.HasSuffix(r.URL.Path, "/") || err != nil {
		http.Error(w, "Page Not Found", 404)
	}
	fs.ServeHTTP(w, r)
}
