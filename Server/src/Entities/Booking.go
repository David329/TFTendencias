package entities

import (
	"gopkg.in/mgo.v2/bson"
)

// Booking Entitie.
type Booking struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       string
	FlightID     string
	PersonalSeat Seat
}
