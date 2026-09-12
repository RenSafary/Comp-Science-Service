package database

import (
	"database/sql"
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
func (d *Users) CheckIfUserExists(FirstName, LastName, username, password, class string) string {
	var notExists bool

	query := `SELECT NOT EXISTS(SELECT 1 FROM users WHERE first_name = $1 AND last_name = $2 AND class = $1 AND username = $3)`

	err := d.DB.QueryRow(query, FirstName, LastName, username).Scan(&notExists)
	if err != nil {
		log.Println(sql.ErrNoRows.Error())
		return "No rows"
	}

	if notExists {
		query := `INSERT INTO users (class, first_name, last_name, password, admin, username) VALUES($1, $2, $3, $4, $5, $6)`

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Ошибка хеширования пароля:", err)
			return "Security error"
		}

		_, err = d.DB.Exec(query, class, FirstName, LastName, hashedPassword, false, username)
		if err != nil {
			log.Println("Error adding a user", err)
			return "Could not add the user"
		}

		return "Success"
	}

	return "User already exists"
}
