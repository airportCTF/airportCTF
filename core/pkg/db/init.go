package db

func (pg *Postgres) init() error {
	// Create user table
	createUserTable := `
    CREATE TABLE IF NOT EXISTS users (
        login TEXT PRIMARY KEY,
        password_hash TEXT NOT NULL,
        is_admin BOOLEAN NOT NULL,
        passport_num TEXT NOT NULL,
        name TEXT NOT NULL,
        last_name TEXT NOT NULL
    );
    `

	// Create ticket table
	createTicketTable := `
    CREATE TABLE IF NOT EXISTS tickets (
        pnr TEXT PRIMARY KEY,
        booking_number TEXT NOT NULL,
        passport_num TEXT NOT NULL,
        flight_number TEXT NOT NULL,
        datetime TEXT NOT NULL
    );
    `

	// Create flight table
	createFlightTable := `
    CREATE TABLE IF NOT EXISTS flights (
        number TEXT PRIMARY KEY,
        air_company TEXT NOT NULL,
        from_airport TEXT NOT NULL,
        to_airport TEXT NOT NULL,
        date TEXT NOT NULL,
        plane TEXT NOT NULL
    );
    `

	_, err := pg.db.Exec(createUserTable)
	if err != nil {
		return err
	}
	_, err = pg.db.Exec(createTicketTable)
	if err != nil {
		return err
	}
	_, err = pg.db.Exec(createFlightTable)
	if err != nil {
		return err
	}
	return nil
}
