package flight

import (
	"database/sql"
	"github.com/chessnok/airportCTF/core/pkg/flight"
)

type Flights struct {
	db *sql.DB
}

func NewFlights(db *sql.DB) *Flights {
	return &Flights{db: db}
}

func (f Flights) PutToDB(flight *flight.Flight) error {
	_, err := f.db.Exec("INSERT INTO flights (id, from, to, datetime) VALUES ($1, $2, $3, $4)", flight.ID, flight.From, flight.To, flight.Date)
	return err
}
