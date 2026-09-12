package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvKey() (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
		return "", nil
	}
	key := os.Getenv("SessionKey")

	return key, nil
}

func EncryptSession(username string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}
}

func DecryptSession() (string, error) {

}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   1800, // 30 mins
	}
}
