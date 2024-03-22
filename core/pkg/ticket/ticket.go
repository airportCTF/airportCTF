package ticket

import (
	"github.com/chessnok/airportCTF/core/pkg/flight"
	"math/rand"
	"time"
)

// Ticket is a pass to the Flight that user can buy
type Ticket struct {
	PNR            int // random id of ticket, 10-digit number can contain only numbers
	PassportNumber int // passport number that contains 9 digits
	Flight         flight.Flight
	Datetime       time.Time
}

func CreateTicket() *Ticket {
	return &Ticket{
		PNR:            rangeIn(1000000000, 9999999999),
		PassportNumber: 0,
		Flight:         flight.Flight{},
		Datetime:       time.Time{},
	}
}

func rangeIn(low, hi int) int {
	return low + rand.Intn(hi-low)
}
