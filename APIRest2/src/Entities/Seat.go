package entities

// Seat Entitie.
type Seat struct {
	UserID string `bson:"userid" json:"userid"`
	Number string `bson:"number" json:"number"`
	Type   string `bson:"type" json:"type"`
}
