package entities

import (
	"gopkg.in/mgo.v2/bson"
)

// Booking Entitie.
type Booking struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	UserID       string        `bson:"userid" json:"userid"`
	FlightID     string        `bson:"flightid" json:"flightid"`
	PersonalSeat Seat          `bson:"personalseat" json:"personalseat"`
}
