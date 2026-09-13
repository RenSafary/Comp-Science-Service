package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetCryptoKey() ([]byte, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
		return []byte{}, err
	}
	keyStr := os.Getenv("CryptoKey")
	keyBytes := []byte(keyStr)

	return keyBytes, nil
}

func EncryptSession(username string) (string, error) {
	encryptionKey, err := GetCryptoKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(username), nil)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
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
