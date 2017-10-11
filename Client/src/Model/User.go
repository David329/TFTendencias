package model

// User Entitie.
type User struct {
	FirstName      string
	LastName       string
	PassportType   string
	PassportNumber string
	Email          string
	Password       string
	PersonalCard   Payment
}
