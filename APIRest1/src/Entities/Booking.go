package entities

import (
	"gopkg.in/mgo.v2/bson"
)

// Booking Entitie.
type Booking struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	UserID       string        `bson:"userid"`
	FlightID     string        `bson:"flightid"`
	PersonalSeat Seat          `bson:"personalseat"`
}
