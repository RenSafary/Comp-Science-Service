package database

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	DB *sql.DB
}

func UsersInit(db *sql.DB) *Users {
	return &Users{DB: db}
}

/*
func (d *Users) CheckUserPass(username, password string) error {

}
*/

// Sign up
func (d *Users) CreateUser(FirstName, LastName, username, password, class string) string {
	var notExists bool

	query := `SELECT NOT EXISTS(SELECT 1 FROM users WHERE first_name = $1 AND last_name = $2 AND class = $3 AND username = $4)`

	err := d.DB.QueryRow(query, FirstName, LastName, class, username).Scan(&notExists)
	if err != nil {
		log.Println("Database query error:", err)
		return "database error"
	}

	if notExists {
		query := `INSERT INTO users (class, first_name, last_name, password, admin, username) VALUES($1, $2, $3, $4, $5, $6)`

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error hashing pasword:", err)
			return "security error"
		}

		_, err = d.DB.Exec(query, class, FirstName, LastName, hashedPassword, false, username)
		if err != nil {
			log.Println("Error adding a user", err)
			return "could not add the user"
		}

		return "success"
	}

	return "user already exists"
}

func (d *Users) VerifyPassword(username, password string) (bool, error) {
	var hashedPassword string

	query := `SELECT password FROM users WHERE username = $1`
	err := d.DB.QueryRow(query, username).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, errors.New("user not found")
		}
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
