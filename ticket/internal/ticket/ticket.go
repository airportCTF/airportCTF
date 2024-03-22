package ticket

import (
	"github.com/chessnok/airportCTF/core/pkg/flight"
	"github.com/chessnok/airportCTF/core/pkg/ticket"
	"math/rand"
	"time"
)

func CreateTicket() *ticket.Ticket {
	return &ticket.Ticket{
		PNR:            0,
		PassportNumber: 0,
		Flight:         flight.Flight{},
		Datetime:       time.Time{},
	}
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
