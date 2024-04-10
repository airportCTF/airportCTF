package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type requestInfo struct {
	Passengers []string `json:"passengers" xml:"passengers"` // an array of passengers that's registered for that flight
	FlightId   string   `json:"flight_id" xml:"flight_id"`   // every flight has its own id
}

// randBookingNumberGen - generates random string in specific length
func genRandString(length int, lowercase bool) string {
	var out string
	var value int
	switch {
	case lowercase:
		value = 'a'
	default:
		value = 'A'
	}
	for i := 0; i < length; i++ {
		out += string(byte(rand.Intn(26) + value))
	}
	return out
}

func NewBooking(db *db.Postgres, logger *log.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(requestInfo)
		if err := c.Bind(req); err != nil {
			logger.Printf("[ERROR] NewBooking: cant validate request data: %v", c.Request())
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "server error: can not validate request data",
			})
		}
		if len(req.Passengers) == 0 {
			logger.Printf("[WARNING] NewBooking: no passengers specified, request passengers array length: %v", len(req.Passengers))
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "no passengers specified",
			})
		}
		flight, err := db.Flights.GetFromDB(req.FlightId)
		if err != nil {
			logger.Printf("[ERROR] NewBooking: cannot get flight from DB. request flightID info: %v", req.FlightId)
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Internal server error",
			})
		}
		if flight == nil {
			logger.Printf("[WARNING] NewBooking: flight not found")
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "flight not found",
			})
		}
		bookingNumber := genRandString(10, false)
		tickets := make([]ticket.Ticket, 0, len(req.Passengers))
		for _, login := range req.Passengers {
			u, err := db.Users.GetFromDB(login)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Internal server error",
				})
			}
			if u == nil {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "user not found",
				})
			}
			passportNumber := u.PassportNum // todo: passport number надо брать из профиля юзера, так как юзер пост запросом POST v1/ticket покупает билет. бмлет записывается в базу данных. количество билетов регулируется рейсом.
			pnr := genRandString(20, true)
			tickets = append(tickets, ticket.Ticket{
				PNR:            pnr,
				BookingNumber:  bookingNumber,
				PassportNumber: passportNumber,
				Flight:         *flight,
				Datetime:       time.Now(),
			})
		}
		for _, t := range tickets {
			if err := db.Tickets.PutToDB(&t); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Internal server error",
				})
			}
		}
		return c.JSON(http.StatusOK, tickets)
	}
}
