package routers

import (
	"database/sql"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("Could not parse 'sign-in' form")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
}

func SignUp() {

}
