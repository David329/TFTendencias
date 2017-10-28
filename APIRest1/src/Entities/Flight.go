package entities

import "gopkg.in/mgo.v2/bson"

// Flight Entitie.
type Flight struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	AirplaneModel  string        `bson:"airplanemodel"`
	AirplaneNumber string        `bson:"airplanenumber"`
	Price          float32       `bson:"price"`
	Depart         Departure     `bson:"depart"`
	Destin         Destination   `bson:"destin"`
	Seats          []Seat        `bson:"seats"`
}
