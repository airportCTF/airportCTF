package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func NewServer(logger *log.Logger, db *db.Postgres) *echo.Echo {
	server := echo.New()
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		for i := 0; i <= 20; i++ {
			secretKey += strconv.Itoa(rand.Intn(10))
		}
	}

	logginMiddlewar := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ses, err := c.Cookie("session")
			if err != nil {
				c.Set("user", nil)
				return next(c)
			}
			if ses.Value == "" {
				c.Set("user", nil)
				return next(c)
			}
			tokenString := ses.Value
			tokenInfo, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					cookie := &http.Cookie{
						Name:   "session",
						Value:  "",
						MaxAge: -1,
					}
					c.SetCookie(cookie)
				}
				return []byte(secretKey), nil
			})
			if err != nil {
				c.Set("user", nil)
				return next(c)
			}
			if claims, ok := tokenInfo.Claims.(jwt.MapClaims); ok {
				u, err := db.Users.GetFromDB(claims["login"].(string))
				if err != nil {
					cookie := &http.Cookie{
						Name:   "session",
						Value:  "",
						MaxAge: -1,
					}
					c.SetCookie(cookie)
					c.Set("user", nil)
					return next(c)
				}
				c.Set("user", u)
			} else {
				cookie := &http.Cookie{
					Name:   "session",
					Value:  "",
					MaxAge: -1,
				}
				c.SetCookie(cookie)
				c.Set("user", nil)
				return next(c)
			}
			return next(c)
		}
	}
	loggingMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			logger.Println(c.Request().Method, c.Request().URL)
			return next(c)
		}
	}
	server.Use(loggingMiddleware)
	server.Use(logginMiddlewar)
	g := server.Group("/v1")
	g.GET("/profile", GetProfile())
	g.POST("/register", Register(db))
	g.POST("/login", Login(db, secretKey))
	g.GET("/logout", Logout())
	g.POST("/make_admin", MakeAdmin(db))
	return server
}
