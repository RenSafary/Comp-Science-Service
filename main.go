package main

import (
	"Comp-Science-Service/routers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	static := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	r.HandleFunc("/", routers.MainPage).Methods("GET")
	r.HandleFunc("/sign_in", routers.SignIn).Methods("POST", "GET")

	ip := "http://127.0.0.1"
	port := ":8080"

	fmt.Printf("Server started on %s%s", ip, port)
	http.ListenAndServe(port, r)
}
