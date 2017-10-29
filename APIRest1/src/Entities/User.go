package entities

import "gopkg.in/mgo.v2/bson"

// User Entitie.
type User struct {
	ID             bson.ObjectId `bson:"_id,omitempty" json:"_id"` //bson id, omitempty permite vacios en id
	FirstName      string        `bson:"firstname" json:"firstname"`
	LastName       string        `bson:"lastname" json:"lastname"`
	PassportType   string        `bson:"passporttype" json:"passporttype"`
	PassportNumber string        `bson:"passportnumber" json:"passportnumber"`
	Email          string        `bson:"email" json:"email"`
	Password       string        `bson:"password" json:"password"`
	PersonalCard   Payment       `bson:"personalcard" json:"personalcard"`
}
