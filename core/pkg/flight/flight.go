package flight

type Flight struct {
	Number     string
	AirCompany string
	From       string
	To         string
	Date       string
	Plane      string
}

func NewFlight(number, airCompany, from, to, date, plane string) Flight {
	return Flight{
		Number:     number,
		AirCompany: airCompany,
		From:       from,
		To:         to,
		Date:       date,
		Plane:      plane,
	}
}
