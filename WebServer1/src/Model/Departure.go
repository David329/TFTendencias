package model

// Departure Entitie.
type Departure struct {
	Country string `json:"country"`
	City    string `json:"city"`
	TD      string `json:"td"` //departure, tndria q ser un dateTime
	TA      string `json:"ta"` //arrival, tndria q ser un dateTime
}
