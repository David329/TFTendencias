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

//GetAllBooking Envia todos las reservar, formato->JSON
func GetAllBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []interface{}
	DB.GetObjs("Bookings", &obj)

	response(&wr, &obj[0])
}

//GetBookingByID Envia la reserva por ID, formato->JSON
func GetBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{}

	DB.GetObjsByID("Bookings", ps.ByName("id"), &obj)

	response(&wr, &obj)
}

//PostBooking Inserta un nuevo vuelo
func PostBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//obtener el json y lo guardo en body
	var booking interface{} = new(Entities.Booking)
	// var booking Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}
	//parseo de json a Booking, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &booking)

	//Obtenemos el objeto flight de la reserva, para actualizar su asiento del usuario
	// flight := &Entities.Flight{}
	var flight interface{}
	var auxFlight Entities.Flight
	DB.GetObjsByID("Flights", booking.(*Entities.Booking).FlightID, &flight)

	jsonString, _ := json.Marshal(flight)
	json.Unmarshal(jsonString, &auxFlight)

	//chekar si esta en el limite de asientos
	if len(auxFlight.Seats) <= 30 {

		var auxUser Entities.User

		//		auxFlight = flight.(Entities.Flight)

		//agregamos el asiento final a la lista de asientos
		auxFlight.Seats = append(auxFlight.Seats, booking.(*Entities.Booking).PersonalSeat)

		flight = auxFlight
		DB.UpdateObjByID("Flights", booking.(*Entities.Booking).FlightID, &flight)
		//Obtenemos el objeto user de la reserva, para actualizar su monto total
		var user interface{} //= new(Entities.User)
		// user := &Entities.User{}
		DB.GetObjsByID("Users", booking.(*Entities.Booking).UserID, &user)

		jsonString, _ := json.Marshal(user)
		json.Unmarshal(jsonString, &auxUser)
		log.Println("6")
		//restamos el monto total por la reserva
		// auxUser = user.(Entities.User)
		auxUser.PersonalCard.Total -= flight.(Entities.Flight).Price
		log.Println("7")
		user = auxUser
		//actualizamos el monto total del usuario
		DB.UpdateObjByID("Users", booking.(*Entities.Booking).UserID, &user)
		log.Println("8")
		//inserto en la bd de Reservas
		DB.InsertObj("Bookings", &booking)
		log.Println("9")
		//Respuesta
		booking = "Reserva Completada"
		response(&wr, &booking)
	} else {

		//Respuesta
		booking = "Este Vuelo esta lleno, Max 30 Asientos!!!"
		response(&wr, &booking)
	}
}

//PutBookingByID Actualiza un Documento Booking
func PutBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) { //pensar si es correcto...

	var obj interface{} = new(Entities.Booking)
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Bookings", ps.ByName("id"), &obj)

	obj = "ok"
	response(&wr, &obj)
}

//DeleteBookingByID Elimina un usuario por ID, formato->JSON
func DeleteBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{} = "ok"
	DB.DeleteObjByID("Bookings", ps.ByName("id"))

	response(&wr, &obj)
}
