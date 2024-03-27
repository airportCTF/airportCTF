package flight

import "time"

type PlaneStatus int

type AirportCode uint8

const (
	NOTREADY PlaneStatus = iota
	READYTOFLIGHT
	INFLIGHT
)

// Flight is a structure that contains information about which plane will be flying and from where it departure and where it arrives
type Flight struct {
	ID   string    `json:"id"` // Flight ID
	From string    `json:"from"`
	To   string    `json:"to"`
	Date time.Time `json:"date"`
}

type Plane struct {
	PlaneOwner     string      // Which AirCompany Owns this Plane
	PlaneCode      string      // unique code of this plane
	AmountOfPlaces int         // how many passengers can be in plane
	PlaneStatus    PlaneStatus // Status of the plane
	location       AirportCode // Where is located last time (AirportCode)
}
