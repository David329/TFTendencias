//Package routes allow methods for Model Booking
package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"github.com/julienschmidt/httprouter"
)

//GetAllBooking Return All Objects
func GetAllBooking(wr http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	//Return through pointer and save in obj
	var obj []interface{}
	DB.GetObjs("Bookings", &obj)

	response(&wr, &obj[0])
}

//GetBookingByID Return object ByID
func GetBookingByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Return through pointer and save in obj
	var obj interface{}
	DB.GetObjsByID("Bookings", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//PostBooking Insert a new Object
func PostBooking(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//New Obj like Entities.Booking
	var booking interface{} = new(Entities.Booking)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &booking)

	//Get Flight through pointer, New auxFlight to parse from json to Entities.Flight. Objective: Update Seats
	var flight interface{}
	var auxFlight Entities.Flight
	DB.GetObjsByID("Flights", booking.(*Entities.Booking).FlightID, &flight)
	jsonString, _ := json.Marshal(flight)
	json.Unmarshal(jsonString, &auxFlight)

	if len(auxFlight.Seats) <= 30 {

		//Update Seats of the flight
		auxFlight.Seats = append(auxFlight.Seats, booking.(*Entities.Booking).PersonalSeat)

		//Reuse inteface from auxFlight to oldFlight, then update the flight in DB
		flight = auxFlight
		DB.UpdateObjByID("Flights", booking.(*Entities.Booking).FlightID, &flight)

		//Get User through pointer, New auxUser to parse from json to Entities.User. Objective: Update Total Money of User
		var user interface{}
		var auxUser Entities.User
		DB.GetObjsByID("Users", booking.(*Entities.Booking).UserID, &user)
		jsonString, _ := json.Marshal(user)
		json.Unmarshal(jsonString, &auxUser)

		//Update Total mount given price of flight
		auxUser.PersonalCard.Total -= flight.(Entities.Flight).Price

		//Reuse inteface from auxUser to oldUser, then update the User in DB
		user = auxUser
		DB.UpdateObjByID("Users", booking.(*Entities.Booking).UserID, &user)

		//Insert obj, then return through pointer
		DB.InsertObj("Bookings", &booking)

		//Response ok or error
		response(&wr, &booking)
	} else {

		//Response customized or error
		booking = "Este Vuelo esta lleno, Max 30 Asientos!!!"
		response(&wr, &booking)
	}
}

//PutBookingByID Update a Object
func PutBookingByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) { //pensar si es correcto...

	//New Obj like Entities.Booking
	var obj interface{} = new(Entities.Booking)

	//Read Body of Form, then convert json binary to Struct previously defined
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &obj)

	//Update obj, then return through pointer
	DB.UpdateObjByID("Bookings", ps.ByName("id"), &obj)

	//Response ok or error
	response(&wr, &obj)
}

//DeleteBookingByID Delete object by ID
func DeleteBookingByID(wr http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	//Delete object depending Model's ID of url, then return through pointer
	var obj interface{} = ps.ByName("id")
	DB.DeleteObjByID("Bookings", &obj)

	//Response ok or error
	response(&wr, &obj)
}
