package entities

import "gopkg.in/mgo.v2/bson"

// Flight Entitie.
type Flight struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	AirplaneModel  string        `bson:"airplanemodel" json:"airplanemodel"`
	AirplaneNumber string        `bson:"airplanenumber" json:"airplanenumber"`
	Price          float32       `bson:"price" json:"price"`
	Depart         Departure     `bson:"depart" json:"depart"`
	Destin         Destination   `bson:"destin" json:"destin"`
	Seats          []Seat        `bson:"seats" json:"seats"`
}
