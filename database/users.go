package database

import (
	"database/sql"
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
