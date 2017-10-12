package model

// Booking Entitie.
type Booking struct {
	ID             string //en este caso no es bson, xq viene del json no de mgo
	UserID       string
	FlightID     string
	PersonalSeat Seat
}
