package signup

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		http.Error(w, "Must provide an email", http.StatusBadRequest)
		return
	}

	password := r.FormValue("password")
	if password == "" {
		http.Error(w, "Must provide a password", http.StatusBadRequest)
		return
	}

	db, err := sql.Open("mysql", "root:@/user")
	if err != nil {
		http.Error(w, "Sql error", http.StatusInternalServerError)
		return
	}

	_, err = db.Query("INSERT into user (email, password) VALUES ('" + email + "', '" + password + "')")
	if err != nil {
		http.Error(w, "Insert error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}
