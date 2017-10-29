package entities

// Departure Entitie.
type Departure struct {
	Country string `bson:"country" json:"country"`
	City    string `bson:"city" json:"city"`
	TD      string `bson:"td" json:"td"` //departure, tndria q ser un dateTime
	TA      string `bson:"ta" json:"ta"` //arrival, tndria q ser un dateTime
}
