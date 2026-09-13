package auth

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type UserSignUp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Class     string `json:"class"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (s *AuthConf) SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/auth/sign_up.html")
	if err != nil {
		log.Println("Could not parse sign_up.html:", err)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *AuthConf) SignUp(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Sign up websocket err:", err)
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

		result := s.DB.Users.CreateUser(user.FirstName, user.LastName, user.Username, user.Password, user.Class)
		log.Println(result)

		response := []byte(result)

		if err := ws.WriteMessage(msgType, response); err != nil {
			log.Println(err)
			break
		}
	}
}
