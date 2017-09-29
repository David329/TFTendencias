package entities

type Flights struct {
    AirplaneModel string
    AirplaneNumber string
    Price float32
    Depart Departure
    Destin Destination
    Seats[] Seat
}