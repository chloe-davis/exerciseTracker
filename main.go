package main

import (
	"fmt"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/user")
	if err != nil {
		fmt.Println("my sql error", err)
	}

	_, err = db.Query("INSERT into MyGuests (firstname, lastname, email) VALUES ('nick', 'davis', 'nick.davis@sendgrid.com')")
	if err != nil {
		fmt.Println("insert error: ", err)
	}
	fmt.Println("Hello, world.\n")
}
