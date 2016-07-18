package main

import (
	"log"
	"net/http"

	signup "./signup"
	"github.com/gorilla/mux"
)

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":50000", router))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/signup", signup.Signup)

	return router
}
