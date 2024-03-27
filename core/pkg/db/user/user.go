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
	_, err := u.db.Exec("INSERT INTO users (login, password_hash, is_admin, passport_num, name, last_name) VALUES ($1, $2, $3, $4, $5, $6)", user.Login, user.PasswordHash, user.IsAdmin, user.PassportNum, user.Name, user.LastName)
	return err
}
