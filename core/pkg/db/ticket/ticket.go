package ticket

import (
	"database/sql"
	"github.com/chessnok/airportCTF/core/pkg/db/flight"
	ticket2 "github.com/chessnok/airportCTF/core/pkg/ticket"
)

type Tickets struct {
	db *sql.DB
	*flight.Flights
}

func NewTickets(db *sql.DB) *Tickets {
	return &Tickets{db: db, Flights: flight.NewFlights(db)}
}

func (t *Tickets) PutToDB(ticket *ticket2.Ticket) error {
	_, err := t.db.Exec("INSERT INTO tickets (pnr, passport_num, flight_number, datetime) VALUES ($1, $2, $3, $4)", ticket.PNR, ticket.PassportNumber, ticket.Flight.ID, ticket.Datetime)
	return err
}

func (t *Tickets) DeleteFromDB(pnr int) error {
	_, err := t.db.Exec("DELETE FROM tickets WHERE pnr = $1", pnr)
	return err
}

func (t *Tickets) UpdateInDB(ticket *ticket2.Ticket) error {
	_, err := t.db.Exec("UPDATE tickets SET passport_num = $1, flight_number = $2, datetime = $3 WHERE pnr = $4", ticket.PassportNumber, ticket.Flight.ID, ticket.Datetime, ticket.PNR)
	return err
}
func (t *Tickets) GetFromDB(pnr int) (*ticket2.Ticket, error) {
	row := t.db.QueryRow("SELECT pnr, passport_num, flight_number, datetime FROM tickets WHERE pnr = $1", pnr)
	tick := &ticket2.Ticket{}
	err := row.Scan(&tick.PNR, &tick.PassportNumber, &tick.Flight.ID, &tick.Datetime)
	if err != nil {
		return nil, err
	}
	return tick, nil
}

func (t *Tickets) GetAllFromDB() ([]*ticket2.Ticket, error) {
	rows, err := t.db.Query("SELECT pnr, passport_num, flight_number, datetime FROM tickets")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tickets := make([]*ticket2.Ticket, 0)
	for rows.Next() {
		tick := &ticket2.Ticket{}
		err := rows.Scan(&tick.PNR, &tick.PassportNumber, &tick.Flight.ID, &tick.Datetime)
		if err != nil {
			return nil, err
		}
		fl, erro := t.Flights.GetFromDB(tick.Flight.ID)
		if erro != nil {
			continue
		}
		tick.Flight = *fl
		tickets = append(tickets, tick)
	}
	return tickets, nil
}
