package routers

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Acces-Control-Allow-Headers", "Content-Type")

	tmpl, err := template.ParseFiles("./templates/main_page.html")
	if err != nil {
		log.Println("Couldn't parse main_page.html:", err)
		return
	}

	_, err = r.Cookie("session_token") // _ instead of 'cookie'
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			http.Redirect(w, r, "/sign-in", http.StatusSeeOther)
			return
		}
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
