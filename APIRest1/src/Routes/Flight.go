package routes

//Restful - Flight
import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	DB "../DB"
	Entities "../Entities"

	"github.com/julienschmidt/httprouter"
)

//GetAllFlight Envia todos los vuelos, formato->JSON
func GetAllFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj []interface{}
	DB.GetObjs("Flights", &obj)

	response(&wr, &obj[0])
}

//GetFlightByID Envia El vuelo por ID, formato->JSON
func GetFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{}

	DB.GetObjsByID("Flights", ps.ByName("id"), &obj)
	// jsonString, _ := json.Marshal(obj)
	// log.Println(string(jsonString))
	response(&wr, &obj)
}

//PostFlight Inserta un nuevo vuelo
func PostFlight(wr http.ResponseWriter, req *http.Request, _ httprouter.Params) {

	var obj interface{} = new(Entities.Flight)
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &obj)
	DB.InsertObj("Flights", &obj)

	obj = "ok"
	response(&wr, &obj)
}

//PutFlightByID Actualiza un Documento Flight
func PutFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{} = new(Entities.Flight)
	body, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(body, &obj)
	DB.UpdateObjByID("Flights", ps.ByName("id"), &obj)

	obj = "ok"
	response(&wr, &obj)
}

//DeleteFlightByID Elimina un usuario por ID, formato->JSON
func DeleteFlightByID(wr http.ResponseWriter, req *http.Request, ps httprouter.Params) {

	var obj interface{} = "ok"
	DB.DeleteObjByID("Flights", ps.ByName("id"))

	response(&wr, &obj)
}
