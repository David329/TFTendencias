package model

// Flight Entitie.
type Flight struct {
	ID             string      `json:"_id"`
	AirplaneModel  string      `json:"airplanemodel"`
	AirplaneNumber string      `json:"airplanenumber"`
	Price          float32     `json:"price"`
	Depart         Departure   `json:"depart"`
	Destin         Destination `json:"destin"`
	Seats          []Seat      `json:"seats"`
}
