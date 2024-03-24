package ticket

import (
	"database/sql"
	ticket2 "github.com/chessnok/airportCTF/core/pkg/ticket"
)

type Tickets struct {
	db *sql.DB
}

func NewTickets(db *sql.DB) *Tickets {
	return &Tickets{db: db}
}

func (t *Tickets) PutToDB(ticket *ticket2.Ticket) error {
	_, err := t.db.Exec("INSERT INTO tickets (pnr, passport_number, flight_id, datetime) VALUES ($1, $2, $3, $4)", ticket.PNR, ticket.PassportNumber, ticket.Flight.ID, ticket.Datetime)
	return err
}

func (t *Tickets) GetFromDB(pnr int) (*ticket2.Ticket, error) {
	row := t.db.QueryRow("SELECT pnr, passport_number, flight_id, datetime FROM tickets WHERE pnr = $1", pnr)
	tick := &ticket2.Ticket{}
	err := row.Scan(&tick.PNR, &tick.PassportNumber, &tick.Flight.ID, &tick.Datetime)
	if err != nil {
		return nil, err
	}
	return tick, nil
}
