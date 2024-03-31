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
	_, err := f.db.Exec("INSERT INTO flights (id, from, to, date) VALUES ($1, $2, $3, $4)", flight.ID, flight.From, flight.To, flight.Date)
	return err
}

// GetFromDB - get information about initial flight from database by flight id
func (f Flights) GetFromDB(id string) (*flight.Flight, error) {
	row := f.db.QueryRow("SELECT from_airport, to_airport, date FROM flights WHERE number = $1", id)
	fl := &flight.Flight{ID: id}
	err := row.Scan(&fl.From, &fl.To, &fl.Date)
	if err != nil {
		return nil, err
	}
	return fl, nil
}
