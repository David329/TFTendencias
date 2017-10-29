package model

// User Entitie.
type User struct {
	ID             string  `json:"_id"`
	FirstName      string  `json:"firstname"`
	LastName       string  `json:"lastname"`
	PassportType   string  `json:"passporttype"`
	PassportNumber string  `json:"passportnumber"`
	Email          string  `json:"email"`
	Password       string  `json:"password"`
	PersonalCard   Payment `json:"personalcard"`
}
