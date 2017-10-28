package routes

//Restful - Booking
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"
	"github.com/julienschmidt/httprouter"
)

//Get-Post-Put-Delete

//GetAllBooking Envia todos las reservar, formato->JSON
func GetAllBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todas las reservas por el metodo Generico
	var bookings []Entities.Booking
	bookings = *DB.GetObjs("Bookings", Entities.Booking{}).(*[]Entities.Booking)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(bookings)
}

//GetBookingByID Envia la reserva por ID, formato->JSON
func GetBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	booking := &Entities.Booking{}

	DB.GetObjsByID("Bookings", ps.ByName("id"), &booking)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(booking)
}

//PostBooking Inserta un nuevo vuelo
func PostBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//obtener el json y lo guardo en body
	var booking Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a Booking, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &booking)

	//Obtenemos el objeto flight de la reserva, para actualizar su asiento del usuario
	flight := &Entities.Flight{}
	DB.GetObjsByID("Flights", booking.FlightID, &flight)

	wr.Header().Set("Content-Type", "application/json")

	//chekar si esta en el limite de asientos
	if len(flight.Seats) <= 30 {

		//agregamos el asiento final a la lista de asientos
		flight.Seats = append(flight.Seats, booking.PersonalSeat)

		DB.UpdateObjByID("Flights", booking.FlightID, flight)
		log.Println(booking.UserID)
		//Obtenemos el objeto user de la reserva, para actualizar su monto total
		user := &Entities.User{}
		DB.GetObjsByID("Users", booking.UserID, &user)

		//restamos el monto total por la reserva
		user.PersonalCard.Total -= flight.Price

		//actualizamos el monto total del usuario
		DB.UpdateObjByID("Users", booking.UserID, user)

		//inserto en la bd de Reservas
		DB.InsertObj("Bookings", booking)

		//Respuesta
		json.NewEncoder(wr).Encode("Reserva Completada")
	} else {

		//Respuesta
		json.NewEncoder(wr).Encode("Este Vuelo esta lleno, Max 30 Asientos!!!")
	}
}

//PutBookingByID Actualiza un Documento Booking
func PutBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) { //pensar si es correcto...

	//obtener el json y lo guardo en body
	var obj Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a Booking
	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Bookings", ps.ByName("id"), obj)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteBookingByID Elimina un usuario por ID, formato->JSON
func DeleteBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	DB.DeleteObjByID("Bookings", ps.ByName("id"))

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Eliminado")
}
