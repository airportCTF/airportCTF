package http

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/chessnok/airportCTF/core/pkg/db"
	user2 "github.com/chessnok/airportCTF/core/pkg/user"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"strconv"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Register(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(RegisterRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if req.Login == "" || req.Password == "" {
			return c.JSON(http.StatusBadRequest, "password or login not provided")
		}
		u, err := db.Users.GetFromDB(req.Login)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if u != nil {
			return c.JSON(http.StatusConflict, "user already exists")
		}
		randomNumber := strconv.Itoa(rand.Intn(9000000000) + 1000000000)
		firstName := randomdata.FirstName(randomdata.RandomGender)
		lastName := randomdata.LastName()
		user := user2.NewUser(req.Login, req.Password, randomNumber, firstName, lastName, false)
		err = db.Users.PutToDB(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusCreated, user)
	}
}
