package auth

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *AuthConf) SignInPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/auth/sign_in.html")
	if err != nil {
		log.Println("Could not parse sign_in.html:", err)
		return
	}

	tmpl.Execute(w, nil)
}

func (s *AuthConf) SignIn(w http.ResponseWriter, r *http.Request) {
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

		var user User
		err = json.Unmarshal(msg, &user)
		if err != nil {
			log.Println("Invalid JSON format:", err)
			continue
		}

		log.Println(user.Username, user.Password)

		response := []byte("success")
		if err := ws.WriteMessage(msgType, response); err != nil {
			log.Println(err)
			break
		}
	}
}
