package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type CookieReq struct {
	Username string `json:"username"`
}

func GetCryptoKey() ([]byte, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	keyStr := os.Getenv("CryptoKey")
	if keyStr == "" {
		return nil, errors.New("CryptoKey is not set in environment")
	}

	keyBytes, err := hex.DecodeString(keyStr)
	if err != nil {
		return nil, err
	}

	if len(keyBytes) != 32 {
		return nil, errors.New("invalid key size: must be exactly 32 bytes (64 hex characters)")
	}

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

/*
func DecryptSession() (string, error) {

}
*/

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CookieReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Username == "" {
		log.Println(err)
		return
	}

	token, err := EncryptSession(req.Username)
	if err != nil {
		log.Println(err)
		return
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   1800, // 30 mins
	}

	http.SetCookie(w, &cookie)
}
