package auth

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type UserSignUp struct {
	Class     string
	FirstName string
	LastName  string
	Username  string
	Password  string
}

func (s *AuthConf) SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/auth/sign_up.html")
	if err != nil {
		log.Println("Could not parse sign_in.html:", err)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *AuthConf) SignUp(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Sign-In websocket err:", err)
		return
	}
	defer ws.Close()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("Received: %s", msg)

		var user UserSignUp
		err = json.Unmarshal(msg, &user)
		if err != nil {
			log.Println("Invalid JSON format:", err)
			continue
		}

		log.Println(user.Class, user.FirstName, user.LastName, user.Password)
		result := s.DB.Users.CheckIfUserExists(user.FirstName, user.LastName, user.Username, user.Password, user.Class)
		if result == "Success" {
			response := []byte("success")
			if err := ws.WriteMessage(msgType, response); err != nil {
				log.Println(err)
				break
			}
		}
	}
}
