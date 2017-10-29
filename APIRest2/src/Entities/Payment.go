package entities

// Payment Entitie.
type Payment struct {
	Card           string  `bson:"card" json:"card"`
	CardNumber     string  `bson:"cardnumber" json:"cardnumber"`
	CSC            string  `bson:"csc" json:"csc"`
	Total          float32 `bson:"total" json:"total"`
	ExpirationDate string  `bson:"expirationdate" json:"expirationdate"` //format=(dd/mm/yyyy)
}
