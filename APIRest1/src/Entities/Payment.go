package entities

// Payment Entitie.
type Payment struct {
	Card           string  `bson:"card"`
	CardNumber     string  `bson:"cardnumber"`
	CSC            string  `bson:"csc"`
	Total          float32 `bson:"total"`
	ExpirationDate string  `bson:"expirationdate"` //format=(dd/mm/yyyy)
}
