package entities

// Seat Entitie.
type Seat struct {
	UserID string `bson:"userid"`
	Number string `bson:"number"`
	Type   string `bson:"firstname"`
}
