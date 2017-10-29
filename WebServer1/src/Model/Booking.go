package model

// Booking Entitie.
type Booking struct {
	ID           string `json:"_id"`
	UserID       string `json:"userid"`
	FlightID     string `json:"flightid"`
	PersonalSeat Seat   `json:"personalseat"`
}
