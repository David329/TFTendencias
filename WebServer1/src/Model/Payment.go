package model

// Payment Entitie.
type Payment struct {
	Card           string  `json:"card"`
	CardNumber     string  `json:"cardnumber"`
	CSC            string  `json:"csc"`
	Total          float32 `json:"total"`
	ExpirationDate string  `json:"expirationdate"` //format=(dd/mm/yyyy)
}
