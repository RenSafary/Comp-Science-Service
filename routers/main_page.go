package routers

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Acces-Control-Allow-Headers", "Content-Type")

	tmpl, err := template.ParseFiles("./templates/main.html")
	if err != nil {
		log.Println("Couldn't parse main_page.html:", err)
		return
	}

	tmpl.Execute(w, nil)
}
