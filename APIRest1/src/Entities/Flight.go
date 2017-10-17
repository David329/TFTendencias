package entities

import "gopkg.in/mgo.v2/bson"

// Flight Entitie.
type Flight struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	AirplaneModel  string
	AirplaneNumber string
	Price          float32
	Depart         Departure
	Destin         Destination
	Seats          []Seat
}
