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
func PostBooking(wr http.ResponseWriter, req *http.Request) { //pensar como haremos la reserva para actualizar el asiento en flight.

	//chekar si el array de asiento de flight es menoroigual de 30
	//actualizar flight en asiento
	//actualizar el payment del usuario
	session := DB.GetDbSession()

	//obtener el json y lo guardo en body
	var obj Entities.Booking
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a Booking, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//inserto en la bd
	c := session.DB("lushflydb").C("Bookings")
	err = c.Insert(obj)
	if err != nil {
		log.Fatal(err)
	}

	DB.SendResCloseSession("Objeto Insertado", session, wr)
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
