package ticket

import (
	"github.com/chessnok/airportCTF/core/pkg/flight"
	"time"
)

// Ticket is a pass to the Flight that user can buy
type Ticket struct {
	PNR            string        `json:"pnr"` // random id of ticket, 10-digit number can contain only numbers
	BookingNumber  string        `json:"bookingNumber"`
	PassportNumber string        `json:"passportNumber"` // passport number that contains 9 digits
	Flight         flight.Flight `json:"flight"`
	Datetime       time.Time     `json:"datetime"`
}
