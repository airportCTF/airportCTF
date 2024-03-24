package db

import (
	"database/sql"
	"github.com/chessnok/airportCTF/core/pkg/db/ticket"
)

type Postgres struct {
	config  *Config
	db      *sql.DB
	tickets *ticket.Tickets
}

func NewPostgres(config *Config) *Postgres {
	return &Postgres{config: config, db: nil}
}

func (pg *Postgres) Connect() error {
	db, err := sql.Open("postgres", pg.config.ConnectionUrl())
	if err != nil {
		return err
	}
	pg.db = db
	if err := pg.db.Ping(); err != nil {
		return err
	}
	if err := pg.init(); err != nil {
		return err
	}
	return nil
}