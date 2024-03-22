package user

type User struct {
	ID           int
	Login        string
	PasswordHash string
	IsAdmin      bool
	PassportNum  string
	Name         string
	LastName     string
}

func NewUser(login, passwordHash, passportNum, name, lastName string, isAdmin bool) User {
	return User{
		Login:        login,
		PasswordHash: passwordHash,
		IsAdmin:      isAdmin,
		PassportNum:  passportNum,
		Name:         name,
		LastName:     lastName,
	}
}
