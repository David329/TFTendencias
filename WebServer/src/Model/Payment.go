package model

// Payment Entitie.
type Payment struct {
	Card           string
	CardNumber     string
	CSC            string
	Total          float32
	ExpirationDate string //format=(dd/mm/yyyy)
}
