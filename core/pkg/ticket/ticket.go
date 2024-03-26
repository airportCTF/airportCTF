package ticket

import (
	"github.com/chessnok/airportCTF/core/pkg/flight"
	"time"
)

// Ticket is a pass to the Flight that user can buy
type Ticket struct {
	PNR            int           `json:"pnr"` // random id of ticket, 10-digit number can contain only numbers
	BookingNumber  int           `json:"bookingNumber"`
	PassportNumber int           `json:"passportNumber"` // passport number that contains 9 digits
	Flight         flight.Flight `json:"flight"`
	Datetime       time.Time     `json:"datetime"`
}
