package model

// Flight Entitie.
type Flight struct {
	AirplaneModel  string
	AirplaneNumber string
	Price          float32
	Depart         Departure
	Destin         Destination
	Seats          []Seat
}
