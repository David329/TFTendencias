package routes

//Restful - Booking
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	DB "../DB"
	Entities "../Entities"
	"gopkg.in/mgo.v2/bson"
)

//Get-Post-Put-Delete

//GetAllBooking Envia todos las reservar, formato->JSON
func GetAllBooking(wr http.ResponseWriter, req *http.Request) {

	session := DB.GetDbSession() //en mayusculas pa q sea publico

	//Pa' Obtener
	var bookings []Entities.Booking
	c := session.DB("lushflydb").C("Bookings")
	err := c.Find(nil).Sort("-start").All(&bookings) //es opcional el sort
	if err != nil {
		panic(err)
	}

	//cerrramos sesion
	session.Close()

	//parseamos a json
	data, err := json.Marshal(bookings)
	if err != nil {
		log.Print(err)
		return
	}

	DB.SendResCloseSession(string(data), session, wr)
}

//PostBooking Inserta un nuevo vuelo
func PostBooking(wr http.ResponseWriter, req *http.Request) {
	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var booking Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a Booking, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &booking)

	//Obtenemos el objeto flight de la reserva, para actualizar su asiento del usuario
	var flight Entities.Flight
	c := session.DB("lushflydb").C("Flights")
	err = c.FindId(bson.ObjectIdHex(booking.FlightID)).One(&flight)
	if err != nil {
		panic(err)
	}
	//chekar si esta en el limite de asientos
	if len(flight.Seats) <= 30 {

		//agregamos el asiento final a la lista de asientos
		flight.Seats = append(flight.Seats, booking.PersonalSeat)

		//actualizamos los asientos en el vuelo
		err = c.UpdateId(bson.ObjectIdHex(booking.FlightID), flight)
		if err != nil {
			panic(err)
		}

		//Obtenemos el objeto user de la reserva, para actualizar su monto total
		var user Entities.User
		c = session.DB("lushflydb").C("Users")
		err = c.FindId(bson.ObjectIdHex(booking.UserID)).One(&user)
		if err != nil {
			panic(err)
		}

		//restamos el monto total por la reserva
		user.PersonalCard.Total -= flight.Price

		//actualizamos el monto total del usuario
		err = c.UpdateId(bson.ObjectIdHex(booking.UserID), user)
		if err != nil {
			panic(err)
		}

		//inserto en la bd de Reservas
		c := session.DB("lushflydb").C("Bookings")
		err = c.Insert(booking)
		if err != nil {
			log.Fatal(err)
		}

		DB.SendResCloseSession("Reserva Completada", session, wr)
	} else {
		DB.SendResCloseSession("Este Vuelo esta lleno, Max 30 Asientos!!!", session, wr)
	}
}

//PutBookingByID Actualiza un Documento Booking
func PutBookingByID(wr http.ResponseWriter, req *http.Request) { //pensar si es correcto...

	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a Booking
	json.Unmarshal(body, &obj)

	//obtener el id desde la url
	reqID := req.URL.Query().Get(":id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Bookings")
	err = c.UpdateId(bson.ObjectIdHex(reqID), obj)
	if err != nil {
		panic(err)
	}

	DB.SendResCloseSession("Actualizacion Exitosa", session, wr)
}

//DeleteBookingByID Elimina un usuario por ID, formato->JSON
func DeleteBookingByID(wr http.ResponseWriter, req *http.Request) {

	session := DB.GetDbSession()

	//obtener el id desde la url
	reqID := req.URL.Query().Get(":id")

	//obtener solo los q tienen ese id
	c := session.DB("lushflydb").C("Bookings")
	err := c.RemoveId(bson.ObjectIdHex(reqID))

	if err != nil {
		panic(err)
	}

	DB.SendResCloseSession("Eliminacion Exitosa", session, wr)
}
