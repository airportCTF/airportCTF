package user

import (
	"database/sql"
	"errors"
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

func (u Users) GetFromDB(login string) (*user.User, error) {
	row := u.db.QueryRow("SELECT login, password_hash, is_admin, passport_num, name, last_name FROM users WHERE login = $1", login)
	uu := &user.User{}
	err := row.Scan(&uu.Login, &uu.PasswordHash, &uu.IsAdmin, &uu.PassportNum, &uu.Name, &uu.LastName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return uu, nil
}
