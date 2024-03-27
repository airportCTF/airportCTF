package http

import (
	"github.com/chessnok/airportCTF/core/pkg/db"
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

type requestInfo struct {
	Passengers []string `json:"passengers" xml:"passengers"`
	FlightId   string   `json:"flightId" xml:"flightId"`
}

func NewBooking(db *db.Postgres) func(c echo.Context) error {
	return func(c echo.Context) error {
		req := new(requestInfo)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if len(req.Passengers) == 0 {
			return c.JSON(http.StatusBadRequest, "no passengers specified")
		}
		flight, err := db.Flights.GetFromDB(req.FlightId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if flight == nil {
			return c.JSON(http.StatusNotFound, "flight not found")
		}
		bookingNumber := ""
		for i := 0; i < 10; i++ {
			bookingNumber += string(byte(rand.Intn(26) + 'A'))
		}
		tickets := make([]ticket.Ticket, 0, len(req.Passengers))
		for _, login := range req.Passengers {
			//TODO: заменить на использование логина вместо паспортных данных когда будет готово API авторизации
			pnr := ""
			for i := 0; i < 20; i++ {
				pnr += string(byte(rand.Intn(26) + 'a'))
			}
			tickets = append(tickets, ticket.Ticket{
				PassportNumber: login,
				Flight:         *flight,
				Datetime:       time.Now(),
				BookingNumber:  bookingNumber,
				PNR:            pnr,
			})
		}
		for _, ticket := range tickets {
			if err := db.Tickets.PutToDB(&ticket); err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusOK, tickets)
	}
}
