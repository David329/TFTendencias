package model

// User Entitie.
type User struct {
	ID             string //en este caso no es bson, xq viene del json no de mgo
	FirstName      string
	LastName       string
	PassportType   string
	PassportNumber string
	Email          string
	Password       string
	PersonalCard   Payment
}
