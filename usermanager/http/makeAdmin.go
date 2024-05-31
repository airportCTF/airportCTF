package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/labstack/echo/v4"
	"os"
)

func apply(input string) string {
	var result []rune
	for _, char := range input {
		if char >= '!' && char <= '~' {
			rotated := 33 + ((char - 33 + 47) % 94)
			result = append(result, rotated)
		} else {
			result = append(result, char)
		}
	}
	return string(result)
}
func MakeAdmin(db *db.Postgres) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		apikey := os.Getenv("API_KEY")
		if "Bearer "+apikey != header {
			resp := "Forbidden, expected: Bearer " + apikey[0:len(apikey)-2] + " and two more letters or numbers, but got: " + header
			resp = apply(resp)
			return c.JSON(403, map[string]string{apply("error"): resp})
		}
		user := c.QueryParam("user")
		if user == "" {
			return c.JSON(400, map[string]string{"error": "Bad request"})
		}
		err := db.Users.MakeAdmin(user)
		if err != nil {
			return c.JSON(500, map[string]string{"error": "Internal server error"})
		}
		return c.JSON(200, map[string]string{"status": "Made admin successfully, if user exists"})
	}
}
