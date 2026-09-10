package main

import (
	"Comp-Science-Service/database"
	"Comp-Science-Service/routers"
	"Comp-Science-Service/routers/auth"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db, err := database.Conn()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.DB.Close()

	h := &auth.AuthConf{
		DB: db,
	}

	static := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))

	r.HandleFunc("/", routers.MainPage).Methods("GET")
	r.HandleFunc("/sign_in", h.SignInPage).Methods("GET")
	r.HandleFunc("/sign_in_ws", h.SignIn)

	ip := "http://127.0.0.1"
	port := ":8080"

	fmt.Printf("Server started on %s%s", ip, port)
	http.ListenAndServe(port, r)
}
