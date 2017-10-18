package entities

import "gopkg.in/mgo.v2/bson"

// User Entitie.
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"` //bson id, omitempty permite vacios en id
	FirstName      string
	LastName       string
	PassportType   string
	PassportNumber string
	Email          string
	Password       string
	PersonalCard   Payment
}
