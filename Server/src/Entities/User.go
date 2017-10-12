package entities

import "gopkg.in/mgo.v2/bson"

// User Entitie.
type User struct {
	ID             bson.ObjectId `bson:"_id"`
	FirstName      string
	LastName       string
	PassportType   string
	PassportNumber string
	Email          string
	Password       string
	PersonalCard   Payment
}
