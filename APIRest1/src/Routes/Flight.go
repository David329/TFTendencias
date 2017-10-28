package routes

//Restful - Flight
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

//GetAllFlight Envia todos los vuelos, formato->JSON
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//Obtener Todos los vuelos por el metodo Generico
	var flights []Entities.Flight
	flights = *DB.GetObjs("Flights", Entities.Flight{}).(*[]Entities.Flight)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(flights)
}

//GetFlightByID Envia El vuelo por ID, formato->JSON
func GetFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	flight := &Entities.Flight{}

	DB.GetObjsByID("Flights", ps.ByName("id"), &flight)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode(flight)
}

//PostFlight Inserta un nuevo vuelo
func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.Flight
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a flight, nose si parsea mas de 1 objeto..., seguro con un for o algo
	json.Unmarshal(body, &obj)

	//Insercion en metodo generico
	DB.InsertObj("Flights", obj)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Insertado")
}

//PutFlightByID Actualiza un Documento Flight
func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	//obtener el json y lo guardo en body
	var obj Entities.Flight
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Print(err)
	}

	//parseo de json a flight
	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Flights", ps.ByName("id"), obj)

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Actualizado")
}

//DeleteFlightByID Elimina un usuario por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	DB.DeleteObjByID("Flights", ps.ByName("id"))

	//Respuesta
	wr.Header().Set("Content-Type", "application/json")
	json.NewEncoder(wr).Encode("Objeto Eliminado")
}
