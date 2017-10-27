package entities

import "gopkg.in/mgo.v2/bson"

// User Entitie.
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty"` //bson id, omitempty permite vacios en id
	FirstName      string        `bson:"firstname"`
	LastName       string        `bson:"lastname"`
	PassportType   string        `bson:"passporttype"`
	PassportNumber string        `bson:"passportnumber"`
	Email          string        `bson:"email"`
	Password       string        `bson:"password"`
	PersonalCard   Payment
}
