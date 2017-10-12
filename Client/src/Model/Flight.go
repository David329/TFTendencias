package model

// Flight Entitie.
type Flight struct {
	ID             string //en este caso no es bson, xq viene del json no de mgo
	AirplaneModel  string
	AirplaneNumber string
	Price          float32
	Depart         Departure
	Destin         Destination
	Seats          []Seat
}
