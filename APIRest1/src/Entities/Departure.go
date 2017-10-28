package entities

// Departure Entitie.
type Departure struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	TD      string `bson:"td"` //departure, tndria q ser un dateTime
	TA      string `bson:"ta"` //arrival, tndria q ser un dateTime
}
