package user

import (
	"crypto/sha1"
	"encoding/base64"
)

type User struct {
	Login        string `json:"login" xml:"login"`
	PasswordHash string `json:"passwordHash" xml:"passwordHash"`
	IsAdmin      bool   `json:"isAdmin" xml:"isAdmin"`
	PassportNum  string `json:"passportNum" xml:"passportNum"`
	Name         string `json:"name" xml:"name"`
	LastName     string `json:"lastName" xml:"lastName"`
}

func NewUser(login, password, passportNum, name, lastName string, isAdmin bool) *User {
	u := &User{
		Login:       login,
		IsAdmin:     isAdmin,
		PassportNum: passportNum,
		Name:        name,
		LastName:    lastName,
	}
	u.SetPassword(password)
	return u
}

func (u *User) ComparePassword(password string) bool {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return u.PasswordHash == sha
}

func (u *User) SetPassword(password string) {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	u.PasswordHash = sha
}
