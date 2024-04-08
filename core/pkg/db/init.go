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
        datetime timestamp NOT NULL
    );
    `

	// Create flight table
	createFlightTable := `
    CREATE TABLE IF NOT EXISTS flights (
        id TEXT PRIMARY KEY,
        from_airport TEXT NOT NULL,
        to_airport TEXT NOT NULL,
        datetime TIMESTAMP NOT NULL
    );
    `
	//createPlaneTable := `
	//CREATE TABLE IF NOT EXISTS Planes (
	//    		plane TEXT PRIMARY KEY,
	//    		places INT NOT NULL
	//    	);
	//INSERT INTO Planes (plane, places) VALUES ('Boeing 737', 150);
	//INSERT INTO Planes (plane, places) VALUES ('Boeing 747', 300);
	//INSERT INTO Planes (plane, places) VALUES ('Airbus A320', 180);
	//INSERT INTO Planes (plane, places) VALUES ('Airbus A380', 500);
	//`
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
	//_, err = pg.db.Exec(createPlaneTable)
	//if err != nil {
	//	return err
	//}
	return nil
}
