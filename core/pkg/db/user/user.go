package user

import (
	"database/sql"
	"github.com/chessnok/airportCTF/core/pkg/user"
)

type Users struct {
	db *sql.DB
}

func NewUsers(db *sql.DB) *Users {
	return &Users{db: db}
}

func (u Users) PutToDB(user *user.User) error {
	_, err := u.db.Exec("INSERT INTO users ")
	return err
}
