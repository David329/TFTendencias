package model

// Booking Entitie.
type Booking struct {
	ID           string //en este caso no es bson, xq viene del json no de mgo
	UserID       string //se deberia eliminar, cuando tng tiempo lo pienso
	FlightID     string
	PersonalSeat Seat
}
