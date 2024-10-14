package web_forum

type User struct {
	Found    bool
	Username string
}

func GetUser(Username string) {
	err := Db.QueryRow("SELECT password FROM users WHERE username = ?", Username).Scan(&storedPassword)
}
