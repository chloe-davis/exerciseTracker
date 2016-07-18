package main

import (
	"fmt"
	"log"
	"net/http"

	//"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	//db, err := sql.Open("mysql", "root:@/user")
	//if err != nil {
	//	fmt.Println("my sql error", err)
	//}

	//_, err = db.Query("INSERT into MyGuests (firstname, lastname, email) VALUES ('nick', 'davis', 'nick.davis@sendgrid.com')")
	//if err != nil {
	//	fmt.Println("insert error: ", err)
	//}
	r := mux.NewRouter()
	r.HandleFunc("/signup", Signup)
	log.Fatal(http.ListenAndServe(":50000", r))
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In Signup")
}
